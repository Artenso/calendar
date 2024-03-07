package storage

import (
	"context"
	"time"

	"github.com/Artenso/calendar/internal/model"
)

func (s *Storage) GetFromToEvents(_ context.Context, from time.Time, to time.Time) ([]*model.Event, error) {

	events := make([]*model.Event, 0, 10)
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, storEvent := range s.data {
		if from.Compare(storEvent.Info.EndDate) <= 0 &&
			to.Compare(storEvent.Info.StartDate) >= 0 {
			event := storEvent
			events = append(events, &event)
		}
	}
	return events, nil

}
