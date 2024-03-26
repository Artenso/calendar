package storage

import (
	"context"

	"github.com/Artenso/calendar/internal/model"
	"github.com/Artenso/calendar/internal/storage/calendar/converter"
)

func (s *Storage) GetAllEvents(_ context.Context) ([]*model.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	events := make([]*model.Event, 0, len(s.data))
	for _, storEvent := range s.data {
		event := converter.ToSrvModel(&storEvent)
		events = append(events, event)
	}
	return events, nil
}
