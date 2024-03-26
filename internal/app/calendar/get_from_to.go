package calendar

import (
	"context"

	"github.com/Artenso/calendar/internal/converter"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetFromToEvent gets events from start date to end date
func (i *Implementation) GetFromToEvents(ctx context.Context, req *desc.GetFromToEventsRequest) (*desc.GetFromToEventsResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	utcFrom, utcTo := converter.ToGetFromToEventsDates(req)
	if utcFrom.After(utcTo) {
		return nil, status.Errorf(
			codes.InvalidArgument, "invalid request: %s",
			"wrong from/to format, 'to' must be later than 'from'",
		)
	}

	events, err := i.calendarSrv.GetFromToEvents(ctx, utcFrom, utcTo)
	if err != nil {
		return nil, err
	}

	return converter.ToGetFromToEventsResponse(events), nil
}
