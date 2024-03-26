package calendar

import (
	"context"

	"github.com/Artenso/calendar/internal/converter"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetAllEvents gets events from start date to end date
func (i *Implementation) GetAllEvents(ctx context.Context, _ *emptypb.Empty) (*desc.GetAllEventsResponse, error) {

	events, err := i.calendarSrv.GetAllEvents(ctx)
	if err != nil {
		return nil, err
	}

	return converter.ToGetAllEventsResponse(events), nil
}
