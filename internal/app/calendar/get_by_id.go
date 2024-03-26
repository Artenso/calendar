package calendar

import (
	"context"

	"github.com/Artenso/calendar/internal/converter"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetEventByID gets event from calendar by ID
func (i *Implementation) GetEventByID(ctx context.Context, req *desc.GetEventByIDRequest) (*desc.GetEventByIDResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	event, err := i.calendarSrv.GetEventByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.ToGetEventByIDResponse(event), nil
}
