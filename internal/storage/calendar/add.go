package storage

import (
	"context"
	"time"

	"github.com/Artenso/calendar/internal/model"
	"github.com/Artenso/calendar/internal/storage/calendar/converter"
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

	event := converter.ToRepoModel(info)
	event.ID = eventID
	event.CreatedAt = time.Now().UTC()

	s.data[eventID] = *event

	return converter.ToSrvModel(event), nil
}
