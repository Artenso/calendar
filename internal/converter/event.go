package converter

import (
	"database/sql"
	"time"

	"github.com/Artenso/calendar/internal/model"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToEventInfo(req *desc.AddEventRequest) *model.EventInfo {
	return &model.EventInfo{
		StartDate:   req.GetStartDate().AsTime().UTC(),
		EndDate:     req.GetEndDate().AsTime().UTC(),
		Description: req.GetDescription(),
	}
}

func ToUpdateEventInfo(req *desc.EditEventRequest) *model.UpdateEventInfo {
	return &model.UpdateEventInfo{
		StartDate: sql.NullTime{
			Time:  req.GetStartDate().AsTime().UTC(),
			Valid: req.GetStartDate().IsValid(),
		},
		EndDate: sql.NullTime{
			Time:  req.GetEndDate().AsTime().UTC(),
			Valid: req.GetEndDate().IsValid(),
		},
		Description: sql.NullString{
			String: req.GetDescription().GetValue(),
			Valid:  req.GetDescription() != nil,
		},
	}
}

func toDescEvent(event *model.Event) *desc.Event {
	descEvent := &desc.Event{
		Id:          event.ID,
		StartDate:   timestamppb.New(event.Info.StartDate),
		EndDate:     timestamppb.New(event.Info.EndDate),
		Description: event.Info.Description,
		CreatedAt:   timestamppb.New(event.CreatedAt),
	}
	if event.UpdatedAt.Valid {
		descEvent.UpdatedAt = timestamppb.New(event.UpdatedAt.Time)
	}
	return descEvent
}

func ToAddEventResponse(event *model.Event) *desc.AddEventResponse {
	return &desc.AddEventResponse{
		Event: toDescEvent(event),
	}
}

func ToGetEventByIDResponse(event *model.Event) *desc.GetEventByIDResponse {
	return &desc.GetEventByIDResponse{
		Event: toDescEvent(event),
	}
}

func ToGetFromToEventsDates(req *desc.GetFromToEventsRequest) (time.Time, time.Time) {
	return req.GetFromDate().AsTime().UTC(), req.GetToDate().AsTime().UTC()
}

func ToGetFromToEventsResponse(events []*model.Event) *desc.GetFromToEventsResponse {
	descEventsList := make([]*desc.Event, 0, len(events))
	for _, event := range events {
		descEventsList = append(descEventsList, toDescEvent(event))
	}
	return &desc.GetFromToEventsResponse{
		EventsList: descEventsList,
	}
}

func ToEditEventResponse(event *model.Event) *desc.EditEventResponse {
	return &desc.EditEventResponse{
		NewEvent: toDescEvent(event),
	}
}

func ToGetAllEventsResponse(events []*model.Event) *desc.GetAllEventsResponse {
	descEventsList := make([]*desc.Event, 0, len(events))
	for _, event := range events {
		descEventsList = append(descEventsList, toDescEvent(event))
	}
	return &desc.GetAllEventsResponse{
		EventsList: descEventsList,
	}
}
