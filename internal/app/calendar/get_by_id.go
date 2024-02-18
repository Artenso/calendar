package calendar

import (
	"context"

	"github.com/Artenso/calendar/internal/converter"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
)
func (i *Implementation) GetEventByID(ctx context.Context, req *desc.GetEventByIDRequest) (*desc.GetEventByIDResponse, error){
	event, err := i.calendarSrv.GetEventByID(ctx, req.GetId())

	if err != nil{
		return nil, err
	}

	return converter.ToGetEventByIDResponse(event), nil
}