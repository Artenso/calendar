package service

import (
	"context"
)

// Remove deletes event from calendar
func (s *Service) RemoveEvent(ctx context.Context, eventID int64) error {
	if err := s.eventsStorage.Remove(ctx, eventID); err != nil {
		return err
	}
	return nil
}
