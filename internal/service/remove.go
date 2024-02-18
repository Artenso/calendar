package service

import (
	"context"
)

func (s *service) Remove(ctx context.Context, eventID int64) error {
	err := s.eventsStorage.Remove(ctx, eventID)
	if err != nil {
		return err
	}
	return nil
}
