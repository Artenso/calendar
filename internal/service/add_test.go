package service

import (
	"context"
	"testing"
	"time"

	"github.com/Artenso/calendar/internal/model"
	storageMock "github.com/Artenso/calendar/internal/storage/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAdd(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storage := storageMock.NewMockEventsStorage(ctrl)

	t.Run("OK case", func(t *testing.T) {
		ctx := context.Background()
		in := model.EventInfo{
			StartDate:   time.Date(2023, time.April, 16, 0, 0, 0, 0, time.Now().Location()),
			EndDate:     time.Date(2023, time.April, 16, 23, 59, 59, 0, time.Now().Location()),
			Description: "My Birthday",
		}
		utcIn := model.EventInfo{
			StartDate:   in.StartDate.UTC(),
			EndDate:     in.EndDate.UTC(),
			Description: in.Description,
		}
		mockEvent := model.Event{
			ID:   1,
			Info: utcIn,
		}
		expected := model.Event{
			ID:   1,
			Info: utcIn,
		}
		storage.EXPECT().Add(ctx, utcIn).Return(&mockEvent, nil)
		service := NewService(storage)

		event, err := service.AddEvent(ctx, &in)

		require.NoError(t, err)
		require.Equal(t, &expected, event)
	})

	t.Run("DateBusy error", func(t *testing.T) {
		ctx := context.Background()
		in := model.EventInfo{
			StartDate:   time.Date(2023, time.April, 16, 0, 0, 0, 0, time.Now().Location()),
			EndDate:     time.Date(2023, time.April, 16, 23, 59, 59, 0, time.Now().Location()),
			Description: "My Birthday",
		}
		utcIn := model.EventInfo{
			StartDate:   in.StartDate.UTC(),
			EndDate:     in.EndDate.UTC(),
			Description: in.Description,
		}
		storage.EXPECT().Add(ctx, utcIn).Return(nil, model.ErrDateBusy)
		service := NewService(storage)

		event, err := service.AddEvent(ctx, &in)

		require.Nil(t, event)
		require.EqualError(t, model.ErrDateBusy, err.Error())
	})

	t.Run("StartDate after EndDate error", func(t *testing.T) {
		ctx := context.Background()
		in := model.EventInfo{
			StartDate:   time.Date(2023, time.April, 18, 0, 0, 0, 0, time.Now().Location()),
			EndDate:     time.Date(2023, time.April, 16, 23, 59, 59, 0, time.Now().Location()),
			Description: "My Birthday",
		}

		service := NewService(storage)

		event, err := service.AddEvent(ctx, &in)

		require.Nil(t, event)
		require.Contains(t, err.Error(), "EndDate must be later than StartDate")
	})
}
