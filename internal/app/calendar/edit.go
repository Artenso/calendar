package calendar

import (
	"context"

	"github.com/Artenso/calendar/internal/converter"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EditEvent changes event in calendar
func (i *Implementation) EditEvent(ctx context.Context, req *desc.EditEventRequest) (*desc.EditEventResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	event, err := i.calendarSrv.EditEvent(ctx, req.GetId(), converter.ToUpdateEventInfo(req))
	if err != nil {
		return nil, err
	}

	return converter.ToEditEventResponse(event), nil
}
