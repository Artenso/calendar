package storage

import (
	"context"
	"time"

	"github.com/Artenso/calendar/internal/model"
)

func (s *storage) GetFromToEvents(_ context.Context, from time.Time, to time.Time) []model.Event {

	events := make([]model.Event, 0, 10)
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, event := range s.data {
		if from.Compare(event.Info.EndDate) <= 0 &&
			to.Compare(event.Info.StartDate) >= 0 {
			events = append(events, event)
		}
	}
	return events

}
