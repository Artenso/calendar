package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Artenso/calendar/internal/model"
	storageMock "github.com/Artenso/calendar/internal/storage/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetByID(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storage := storageMock.NewMockEventsStorage(ctrl)

	t.Run("OK case", func(t *testing.T) {
		ctx := context.Background()
		id := int64(25)
		mockEvent := model.Event{
			ID: 25,
			Info: model.EventInfo{
				StartDate:   time.Date(2023, time.April, 16, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2023, time.April, 16, 23, 59, 59, 0, time.UTC),
				Description: "My Birthday",
			},
		}
		storage.EXPECT().GetByID(ctx, id).Return(mockEvent, nil)
		service := NewService(storage)

		event, err := service.GetEventByID(ctx, id)

		require.NoError(t, err)
		require.Equal(t, mockEvent, event)
	})

	t.Run("non-existent event error", func(t *testing.T) {
		ctx := context.Background()
		id := int64(25)

		storage.EXPECT().GetByID(ctx, id).Return(model.Event{}, fmt.Errorf("event not found"))
		service := NewService(storage)

		event, err := service.GetEventByID(ctx, id)

		require.Equal(t, model.Event{}, event)
		require.Error(t, err)
	})
}
