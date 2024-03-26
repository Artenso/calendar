package storage

import (
	"context"

	"github.com/Artenso/calendar/internal/model"
)

func (s *Storage) Remove(_ context.Context, eventID int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[eventID]; ok {
		delete(s.data, eventID)
		return nil
	}
	return model.ErrEventNotFound
}
