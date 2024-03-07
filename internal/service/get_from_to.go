package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Artenso/calendar/internal/model"
)

// GetFromToEvents implements calendar.ICalendarService.
func (s *Service) GetFromToEvents(ctx context.Context, fromDate, toDate time.Time) ([]*model.Event, error) {
	if fromDate.After(toDate) {
		return nil, fmt.Errorf("wrong from/to format, 'to' must be later than 'from'")
	}

	events, err := s.eventsStorage.GetFromToEvents(ctx, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	return events, nil
}
