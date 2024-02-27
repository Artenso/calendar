package converter

import (
	"github.com/Artenso/calendar/internal/model"
	desc "github.com/Artenso/calendar/pkg/calendar/github.com/Artenso/calendar/pkg/calendar"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToEventInfo(req *desc.AddEventRequest) *model.EventInfo {
	return &model.EventInfo{
		StartDate:   req.GetStartTime().AsTime(),
		EndDate:     req.GetEndTime().AsTime(),
		Description: req.GetDescription(),
	}
}

func ToAddEventResponse(event *model.Event) *desc.AddEventResponse {
	return &desc.AddEventResponse{
		Id: event.ID,
	}
}

func ToGetEventByIDResponse(event *model.Event) *desc.GetEventByIDResponse{
	return &desc.GetEventByIDResponse{
		StartTime: timestamppb.New(event.Info.StartDate),
		EndTime: timestamppb.New(event.Info.EndDate),
		Description: event.Info.Description,
	}
}