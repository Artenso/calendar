package calendar

import (
	"context"
	"time"

	"github.com/Artenso/calendar/internal/model"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
)

type ICalendarService interface {
	AddEvent(ctx context.Context, info *model.EventInfo) (*model.Event, error)
	GetEventByID(ctx context.Context, eventID int64) (*model.Event, error)
	RemoveEvent(ctx context.Context, eventID int64) error
	GetFromToEvents(ctx context.Context, fromDate, toDate time.Time) ([]*model.Event, error)
	GetAllEvents(ctx context.Context) ([]*model.Event, error)
	EditEvent(ctx context.Context, id int64, info *model.UpdateEventInfo) (*model.Event, error)
}

type Implementation struct {
	desc.UnimplementedCalendarServer

	calendarSrv ICalendarService
}

func NewCalendar(calendarSrv ICalendarService) *Implementation {
	return &Implementation{
		UnimplementedCalendarServer: desc.UnimplementedCalendarServer{},

		calendarSrv: calendarSrv,
	}
}
