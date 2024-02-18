package service

import (
	"context"
	"fmt"
	"testing"

	storageMock "github.com/Artenso/calendar/internal/storage/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRemove(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storage := storageMock.NewMockEventsStorage(ctrl)

	t.Run("OK case", func(t *testing.T) {

		ctx := context.Background()
		id := int64(25)

		storage.EXPECT().Remove(ctx, id).Return(nil)
		service := NewService(storage)

		err := service.Remove(ctx, id)

		require.NoError(t, err)
	})

	t.Run("non-existent event error", func(t *testing.T) {
		ctx := context.Background()
		id := int64(25)

		storage.EXPECT().Remove(ctx, id).Return(fmt.Errorf("deleting a non-existent event"))
		service := NewService(storage)

		err := service.Remove(ctx, id)

		require.Error(t, err)
	})
}
