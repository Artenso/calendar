package service

import (
	"context"
	"testing"
	"time"

	"github.com/Artenso/calendar/internal/model"
	storageMock "github.com/Artenso/calendar/internal/storage/calendar/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetAll(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storage := storageMock.NewMockIEventsStorage(ctrl)

	t.Run("OK case", func(t *testing.T) {
		ctx := context.Background()

		mockEvents := []model.Event{
			{
				ID: 1,
				Info: model.EventInfo{
					StartDate:   time.Date(2023, time.April, 16, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2023, time.April, 16, 23, 59, 59, 0, time.UTC),
					Description: "My Birthday",
				},
			},
			{
				ID: 2,
				Info: model.EventInfo{
					StartDate:   time.Date(2023, time.November, 2, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2023, time.November, 2, 23, 59, 59, 0, time.UTC),
					Description: "Tony's Birthday",
				},
			},
			{
				ID: 3,
				Info: model.EventInfo{
					StartDate:   time.Date(2023, time.June, 13, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2023, time.June, 13, 23, 59, 59, 0, time.UTC),
					Description: "Dasha's Birthday",
				},
			},
		}

		storage.EXPECT().GetAllEvents(ctx).Return(mockEvents)
		service := NewService(storage)

		events, err := service.GetAllEvents(ctx)

		require.NoError(t, err)
		require.ElementsMatch(t, mockEvents, events)
	})
}
