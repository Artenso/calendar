package service

import (
	"context"
	"fmt"

	"github.com/Artenso/calendar/internal/model"
)

// AddEvent adds event to calendar
func (s *Service) AddEvent(ctx context.Context, info *model.EventInfo) (*model.Event, error) {

	if info.StartDate.After(info.EndDate) {
		return nil, fmt.Errorf("wrong StartDate/EndDate format, EndDate must be later than StartDate")
	}

	event, err := s.eventsStorage.Add(ctx, info)
	if err != nil {
		return nil, err
	}
	return event, nil
}
