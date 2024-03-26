package converter

import (
	srvModel "github.com/Artenso/calendar/internal/model"
	repoModel "github.com/Artenso/calendar/internal/storage/calendar/model"
)

func ToRepoModel(srvInfo *srvModel.EventInfo) *repoModel.Event {
	return &repoModel.Event{
		Info: repoModel.EventInfo{
			StartDate:   srvInfo.StartDate,
			EndDate:     srvInfo.EndDate,
			Description: srvInfo.Description,
		},
	}
}

func ToSrvModel(repoEvent *repoModel.Event) *srvModel.Event {
	srvEvent := &srvModel.Event{
		ID: repoEvent.ID,
		Info: srvModel.EventInfo{
			StartDate:   repoEvent.Info.StartDate,
			EndDate:     repoEvent.Info.EndDate,
			Description: repoEvent.Info.Description,
		},
		CreatedAt: repoEvent.CreatedAt,
	}
	if repoEvent.UpdatedAt.Valid {
		srvEvent.UpdatedAt = repoEvent.UpdatedAt
	}
	return srvEvent
}
