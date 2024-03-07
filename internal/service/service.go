package service

import (
	"context"
	"time"

	"github.com/Artenso/calendar/internal/model"
	"github.com/Artenso/calendar/internal/storage"
)

type IEventsService interface {
	AddEvent(ctx context.Context, event *model.EventInfo) (*model.Event, error)
	RemoveEvent(ctx context.Context, eventID int64) error
	EditEvent(ctx context.Context, eventID int64, info *model.UpdateEventInfo) (*model.Event, error)
	GetAllEvents(ctx context.Context) ([]*model.EventInfo, error)
	GetFromToEvents(ctx context.Context, from, to time.Timer) ([]*model.Event, error)
	GetEventByID(ctx context.Context, eventID int64) (*model.Event, error)
}

type Service struct {
	eventsStorage storage.IEventsStorage
}

func NewService(eventsStorage storage.IEventsStorage) *Service {
	return &Service{
		eventsStorage: eventsStorage,
	}
}
