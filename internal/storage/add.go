package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/Artenso/calendar/internal/model"
)

func (s *Storage) Add(_ context.Context, info *model.EventInfo) (*model.Event, error) {

	s.mu.Lock()
	defer s.mu.Unlock()
	eventID := int64(1)
	if len(s.data) != 0 {
		for id, event := range s.data {
			if info.EndDate.Before(event.Info.StartDate) ||
				info.StartDate.After(event.Info.EndDate) {
				if id > eventID {
					eventID = id
				}
				continue
			}
			return nil, model.ErrDateBusy
		}
		eventID++
	}

	event := model.Event{
		ID:   int64(eventID),
		Info: *info,
		CreatedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}

	s.data[eventID] = event

	return &event, nil
}
