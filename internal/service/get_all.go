package service

import (
	"context"

	"github.com/Artenso/calendar/internal/model"
)

// GetAllEvents gets events from start date to end date
func (s *Service) GetAllEvents(ctx context.Context) ([]*model.Event, error) {
	return s.eventsStorage.GetAllEvents(ctx)
}
