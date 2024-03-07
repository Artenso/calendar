package service

import (
	"context"

	"github.com/Artenso/calendar/internal/model"
)

// GetAllEvents implements calendar.ICalendarService.
func (s *Service) GetAllEvents(ctx context.Context) ([]*model.Event, error) {
	return s.eventsStorage.GetAllEvents(ctx)
}
