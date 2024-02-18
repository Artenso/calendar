package storage

import (
	"context"

	"github.com/Artenso/calendar/internal/model"
)

func (s *storage) Add(_ context.Context, info model.EventInfo) (*model.Event, error) {

	s.mu.Lock()
	defer s.mu.Unlock()
	eventID := int64(0)
	if len(s.data) != 0 {
		for id, item := range s.data {
			if info.EndDate.Before(item.Info.StartDate) ||
				info.StartDate.After(item.Info.EndDate) {
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
		Info: info,
	}

	s.data[eventID] = event

	return &event, nil
}
