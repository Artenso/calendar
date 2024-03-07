package service

import (
	"context"

	"github.com/Artenso/calendar/internal/model"
)

// GetEventByID implements calendar.ICalendarService.
func (s *Service) GetEventByID(ctx context.Context, eventID int64) (*model.Event, error) {

	event, err := s.eventsStorage.GetByID(ctx, eventID)
	if err != nil {
		return nil, err
	}
	return event, nil
}
