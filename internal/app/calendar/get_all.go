package calendar

import (
	"context"

	"github.com/Artenso/calendar/internal/converter"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAllEvents gets events from start date to end date
func (i *Implementation) GetAllEvents(ctx context.Context, req *desc.GetAllEventsRequest) (*desc.GetAllEventsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	events, err := i.calendarSrv.GetAllEvents(ctx)
	if err != nil {
		return nil, err
	}

	return converter.ToGetAllEventsResponse(events), nil
}
