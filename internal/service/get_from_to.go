package service

import (
	"context"
	"time"

	"github.com/Artenso/calendar/internal/model"
)

// GetFromToEvent gets events from start date to end date
func (s *Service) GetFromToEvents(ctx context.Context, fromDate, toDate time.Time) ([]*model.Event, error) {

	events, err := s.eventsStorage.GetFromToEvents(ctx, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	return events, nil
}
