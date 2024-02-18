package storage

import (
	"context"
	"sync"
	"time"

	"github.com/Artenso/calendar/internal/model"
)

type storage struct {
	data map[int64]model.Event
	mu   sync.RWMutex
}

func NewStorage() *storage {
	return &storage{
		data: make(map[int64]model.Event, 0),
	}
}

//go:generate mockgen -source=storage.go -destination=mock/storage_mock.go

type IEventsStorage interface {
	Add(ctx context.Context, info model.EventInfo) (*model.Event, error)
	Remove(ctx context.Context, eventID int64) error
	Edit(ctx context.Context, eventID int64, info model.EventInfo) error
	GetAllEvents(ctx context.Context) []model.Event
	GetFromToEvents(ctx context.Context, from time.Time, to time.Time) []model.Event
	GetByID(ctx context.Context, eventID int64) (*model.Event, error)
}
