package storage

import (
	"context"
	"fmt"
)

func (s *storage) Remove(_ context.Context, eventID int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[eventID]; ok {
		delete(s.data, eventID)
		return nil
	}
	return fmt.Errorf("deleting a non-existent event")
}
