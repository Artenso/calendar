package calendar

import (
	"context"

	"github.com/Artenso/calendar/internal/converter"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddEvent adds event to calendar
func (i *Implementation) AddEvent(ctx context.Context, req *desc.AddEventRequest) (*desc.AddEventResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	event, err := i.calendarSrv.AddEvent(ctx, converter.ToEventInfo(req))
	if err != nil {
		return nil, err
	}

	return converter.ToAddEventResponse(event), nil
}
