package storage

import (
	"context"
	"fmt"

	"github.com/Artenso/calendar/internal/model"
)

func (s *storage) Edit(_ context.Context, eventID int64, info model.EventInfo) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[eventID]; ok {
		s.data[eventID] = model.Event{
			ID:   eventID,
			Info: info,
		}
		return nil
	}
	return fmt.Errorf("changing a non-existent event")
}
