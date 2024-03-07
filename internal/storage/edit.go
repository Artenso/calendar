package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Artenso/calendar/internal/model"
)

func (s *Storage) Edit(_ context.Context, eventID int64, info *model.UpdateEventInfo) (*model.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[eventID]; ok {

		switch {
		case info.EndDate.Valid && info.StartDate.Valid:
			for _, event := range s.data {
				if info.EndDate.Time.Before(event.Info.StartDate) ||
					info.StartDate.Time.After(event.Info.EndDate) {
					continue
				}
				if event.ID != eventID {
					return nil, model.ErrDateBusy
				}
			}
		case info.EndDate.Valid:
			for _, event := range s.data {
				if info.EndDate.Time.Before(event.Info.StartDate) ||
					info.EndDate.Time.After(event.Info.EndDate) {
					continue
				}
				fmt.Printf("event id: %d\nreq id: %d", event.ID, eventID)
				if event.ID != eventID {
					return nil, model.ErrDateBusy
				}
			}
		case info.StartDate.Valid:
			for _, event := range s.data {
				if info.StartDate.Time.After(event.Info.EndDate) || info.StartDate.Time.Before(event.Info.StartDate) {
					continue
				}
				if event.ID != eventID {
					return nil, model.ErrDateBusy
				}
			}
		}

		event := s.data[eventID]
		if info.StartDate.Valid {
			event.Info.StartDate = info.StartDate.Time
		}
		if info.EndDate.Valid {
			event.Info.EndDate = info.EndDate.Time
		}
		if info.Description.Valid {
			event.Info.Description = info.Description.String
		}

		event.UpdatedAt = sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		}

		s.data[eventID] = event

		return &event, nil
	}

	return nil, model.ErrEventNotFound
}
