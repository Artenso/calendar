package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/Artenso/calendar/internal/model"
	"github.com/Artenso/calendar/internal/storage/calendar/converter"
)

func (s *Storage) Edit(_ context.Context, eventID int64, info *model.UpdateEventInfo) (*model.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if updatedEvent, ok := s.data[eventID]; ok {

		if info.StartDate.Valid {
			updatedEvent.Info.StartDate = info.StartDate.Time
		}
		if info.EndDate.Valid {
			updatedEvent.Info.EndDate = info.EndDate.Time
		}
		if info.Description.Valid {
			updatedEvent.Info.Description = info.Description.String
		}

		for _, event := range s.data {
			if updatedEvent.Info.EndDate.Before(event.Info.StartDate) ||
				updatedEvent.Info.StartDate.After(event.Info.EndDate) {
				continue
			}
			if event.ID != eventID {
				return nil, model.ErrDateBusy
			}
		}

		updatedEvent.UpdatedAt = sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		}

		s.data[eventID] = updatedEvent

		return converter.ToSrvModel(&updatedEvent), nil
	}

	return nil, model.ErrEventNotFound
}
