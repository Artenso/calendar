package service

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/Artenso/calendar/internal/model"
	storageMock "github.com/Artenso/calendar/internal/storage/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestEdit(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storage := storageMock.NewMockIEventsStorage(ctrl)

	t.Run("OK case", func(t *testing.T) {
		ctx := context.Background()
		id := int64(25)
		in := model.UpdateEventInfo{
			StartDate: sql.NullTime{
				Time:  time.Date(2023, time.April, 16, 0, 0, 0, 0, time.Now().Location()),
				Valid: true,
			},
			EndDate: sql.NullTime{
				Time:  time.Date(2023, time.April, 16, 23, 59, 59, 0, time.Now().Location()),
				Valid: true,
			},
			Description: sql.NullString{
				String: "My Birthday",
				Valid:  true,
			},
		}
		storage.EXPECT().Edit(ctx, id, in).Return(nil)
		service := NewService(storage)

		_, err := service.EditEvent(ctx, id, &in)

		require.NoError(t, err)
	})

	t.Run("non-existent event error", func(t *testing.T) {
		ctx := context.Background()
		id := int64(25)
		in := model.UpdateEventInfo{
			StartDate: sql.NullTime{
				Time:  time.Date(2023, time.April, 16, 0, 0, 0, 0, time.Now().Location()),
				Valid: true,
			},
			EndDate: sql.NullTime{
				Time:  time.Date(2023, time.April, 16, 23, 59, 59, 0, time.Now().Location()),
				Valid: true,
			},
			Description: sql.NullString{
				String: "My Birthday",
				Valid:  true,
			},
		}
		storage.EXPECT().Edit(ctx, id, in).Return(fmt.Errorf("changing a non-existent event"))
		service := NewService(storage)

		_, err := service.EditEvent(ctx, id, &in)

		require.Error(t, err)
	})
	t.Run("StartDate after EndDate error", func(t *testing.T) {
		ctx := context.Background()
		id := int64(2)
		in := model.UpdateEventInfo{
			StartDate: sql.NullTime{
				Time:  time.Date(2023, time.April, 18, 0, 0, 0, 0, time.Now().Location()),
				Valid: true,
			},
			EndDate: sql.NullTime{
				Time:  time.Date(2023, time.April, 16, 23, 59, 59, 0, time.Now().Location()),
				Valid: true,
			},
			Description: sql.NullString{
				String: "My Birthday",
				Valid:  true,
			},
		}
		service := NewService(storage)

		_, err := service.EditEvent(ctx, id, &in)

		require.Contains(t, err.Error(), "EndDate must be later than StartDate")
	})

}
