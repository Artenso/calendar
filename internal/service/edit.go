package service

import (
	"context"
	"fmt"

	"github.com/Artenso/calendar/internal/model"
)

func (s *service) Edit(ctx context.Context, eventID int64, info model.EventInfo) error {

	if info.StartDate.After(info.EndDate) {
		return fmt.Errorf("wrong StartDate/EndDate format, EndDate must be later than StartDate")
	}

	utcInfo := model.EventInfo{
		StartDate:   info.StartDate.UTC(),
		EndDate:     info.EndDate.UTC(),
		Description: info.Description,
	}

	err := s.eventsStorage.Edit(ctx, eventID, utcInfo)

	if err != nil {
		return err
	}
	return nil
}
