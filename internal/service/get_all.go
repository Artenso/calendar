package service

import (
	"context"

	"github.com/Artenso/calendar/internal/model"
)

func (s *service) GetAllEvents(ctx context.Context) []model.Event {
	return s.eventsStorage.GetAllEvents(ctx)
}
