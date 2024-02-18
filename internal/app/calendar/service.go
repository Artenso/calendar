package calendar

import (
	"context"

	"github.com/Artenso/calendar/internal/model"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
)

type ICalendarService interface{
	AddEvent(ctx context.Context, info *model.EventInfo) (*model.Event, error)
	GetEventByID(ctx context.Context, eventID int64) (*model.Event, error)
}

type Implementation struct {
	desc.UnimplementedCalendarServer

	calendarSrv ICalendarService
}

func NewCalendar(calendarSrv ICalendarService) *Implementation{
	return &Implementation{
		UnimplementedCalendarServer: desc.UnimplementedCalendarServer{},
		
		calendarSrv: calendarSrv,
	}
}
