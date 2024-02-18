package service

import (
	"context"

	"github.com/Artenso/calendar/internal/model"
	"github.com/Artenso/calendar/internal/storage"
)

type IEventsService interface {
	AddEvent(ctx context.Context, event *model.EventInfo) (*model.Event, error)
	Remove(ctx context.Context, eventID int64) error
	Edit(ctx context.Context, eventID int64, info model.EventInfo) error
	GetAllEvents(ctx context.Context) []model.EventInfo
	GetFromToEvents(ctx context.Context, event model.Event) ([]model.Event, error)
	GetEventByID(ctx context.Context, eventID int64) (*model.Event, error)
}

type service struct {
	eventsStorage storage.IEventsStorage
}


func NewService(eventsStorage storage.IEventsStorage) *service {
	return &service{
		eventsStorage: eventsStorage,
	}
}
