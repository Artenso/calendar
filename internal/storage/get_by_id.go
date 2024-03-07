package storage

import (
	"context"

	"github.com/Artenso/calendar/internal/model"
)

func (s *Storage) GetByID(_ context.Context, eventID int64) (*model.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if event, ok := s.data[eventID]; ok {
		return &event, nil
	}
	return nil, model.ErrEventNotFound
}
