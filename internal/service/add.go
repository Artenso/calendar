package service

import (
	"context"
	"fmt"

	"github.com/Artenso/calendar/internal/model"
)

func (s *service) AddEvent(ctx context.Context, info *model.EventInfo) (*model.Event, error) {

	if info.StartDate.After(info.EndDate) {
		return nil, fmt.Errorf("wrong StartDate/EndDate format, EndDate must be later than StartDate")
	}

	utcInfo := model.EventInfo{
		StartDate:   info.StartDate.UTC(),
		EndDate:     info.EndDate.UTC(),
		Description: info.Description,
	}
	event, err := s.eventsStorage.Add(ctx, utcInfo)
	if err != nil {
		return nil, err
	}
	return event, nil
}
