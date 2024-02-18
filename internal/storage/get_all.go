package storage

import (
	"context"

	"github.com/Artenso/calendar/internal/model"
)

func (s *storage) GetAllEvents(_ context.Context) []model.Event {
	s.mu.RLock()
	defer s.mu.RUnlock()

	events := make([]model.Event, 0, len(s.data))
	for _, event := range s.data {
		events = append(events, event)
	}
	return events
}
