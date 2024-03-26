package service

import (
	"context"
	"fmt"

	"github.com/Artenso/calendar/internal/model"
)

// EditEvent changes event in calendar
func (s *Service) EditEvent(ctx context.Context, id int64, info *model.UpdateEventInfo) (*model.Event, error) {

	if info.StartDate.Valid && info.EndDate.Valid {
		if info.StartDate.Time.After(info.EndDate.Time) {
			return nil, fmt.Errorf("wrong StartDate/EndDate format, EndDate must be later than StartDate")
		}
	}

	event, err := s.eventsStorage.Edit(ctx, id, info)
	if err != nil {
		return nil, err
	}
	return event, nil
}
