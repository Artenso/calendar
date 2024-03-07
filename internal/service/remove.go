package service

import (
	"context"
)

// RemoveEvent implements calendar.ICalendarService.
func (s *Service) RemoveEvent(ctx context.Context, eventID int64) error {
	err := s.eventsStorage.Remove(ctx, eventID)
	if err != nil {
		return err
	}
	return nil
}
