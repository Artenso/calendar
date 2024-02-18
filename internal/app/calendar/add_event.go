package calendar

import (
	"context"

	"github.com/Artenso/calendar/internal/converter"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
)
func (i *Implementation) AddEvent(ctx context.Context, req *desc.AddEventRequest) (*desc.AddEventResponse, error){
	event, err := i.calendarSrv.AddEvent(ctx, converter.ToEventInfo(req))

	if err != nil{
		return nil, err
	}

	return converter.ToAddEventResponse(event), nil
}