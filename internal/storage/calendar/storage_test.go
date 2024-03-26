package storage

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Artenso/calendar/internal/model"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {

	storage := New()
	ctx := context.Background()

	t.Run("OK_Add", func(t *testing.T) {

		testTable := []struct {
			name          string
			eventInfo     model.EventInfo
			evetID        int64
			expectedError error
			expectedData  model.Event
		}{
			{
				name: "add_1",
				eventInfo: model.EventInfo{
					StartDate:   time.Date(2023, time.April, 16, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2023, time.April, 16, 23, 59, 59, 0, time.UTC),
					Description: "My Birthday",
				},
				expectedError: nil,
				expectedData: model.Event{
					ID: 0,
					Info: model.EventInfo{
						StartDate:   time.Date(2023, time.April, 16, 0, 0, 0, 0, time.UTC),
						EndDate:     time.Date(2023, time.April, 16, 23, 59, 59, 0, time.UTC),
						Description: "My Birthday",
					},
				},
			},
			{
				name: "add_2",
				eventInfo: model.EventInfo{
					StartDate:   time.Date(2023, time.November, 2, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2023, time.November, 2, 23, 59, 59, 0, time.UTC),
					Description: "Tony's Birthday",
				},
				expectedError: nil,
				expectedData: model.Event{
					ID: 1,
					Info: model.EventInfo{
						StartDate:   time.Date(2023, time.November, 2, 0, 0, 0, 0, time.UTC),
						EndDate:     time.Date(2023, time.November, 2, 23, 59, 59, 0, time.UTC),
						Description: "Tony's Birthday",
					},
				},
			},
			{
				name: "add_3",
				eventInfo: model.EventInfo{
					StartDate:   time.Date(2023, time.June, 13, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2023, time.June, 13, 23, 59, 59, 0, time.UTC),
					Description: "Dasha's Birthday",
				},
				expectedError: nil,
				expectedData: model.Event{
					ID: 2,
					Info: model.EventInfo{
						StartDate:   time.Date(2023, time.June, 13, 0, 0, 0, 0, time.UTC),
						EndDate:     time.Date(2023, time.June, 13, 23, 59, 59, 0, time.UTC),
						Description: "Dasha's Birthday",
					},
				},
			},
		}

		for _, test := range testTable {
			t.Run(test.name, func(t *testing.T) {

				event, err := storage.Add(ctx, &test.eventInfo)

				require.NoError(t, err)
				require.Equal(t, &test.expectedData, event)
			})
		}
	})

	t.Run("BAD_Add", func(t *testing.T) {
		info := model.EventInfo{
			StartDate:   time.Date(2023, time.April, 15, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2023, time.June, 16, 0, 0, 0, 0, time.UTC),
			Description: "Sergey's Birthday",
		}

		event, err := storage.Add(ctx, &info)

		require.EqualError(t, model.ErrDateBusy, err.Error())
		require.Equal(t, (*model.Event)(nil), event)
	})

	t.Run("OK_Edit", func(t *testing.T) {
		id := int64(0)
		info := &model.UpdateEventInfo{
			StartDate: sql.NullTime{
				Time:  time.Date(2023, time.April, 15, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
			EndDate: sql.NullTime{
				Time:  time.Date(2023, time.June, 15, 23, 59, 59, 0, time.UTC),
				Valid: true,
			},
			Description: sql.NullString{
				String: "Sergey's Birthday",
				Valid:  true,
			},
		}

		_, err := storage.Edit(ctx, id, info)
		require.NoError(t, err)
	})

	t.Run("BAD_Edit", func(t *testing.T) {
		id := int64(100)
		info := &model.UpdateEventInfo{
			StartDate: sql.NullTime{
				Time:  time.Date(2023, time.April, 15, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
			EndDate: sql.NullTime{
				Time:  time.Date(2023, time.June, 15, 23, 59, 59, 0, time.UTC),
				Valid: true,
			},
			Description: sql.NullString{
				String: "Sergey's Birthday",
				Valid:  true,
			},
		}
		_, err := storage.Edit(ctx, id, info)
		require.EqualError(t, model.ErrEventNotFound, err.Error())
	})

	t.Run("GetAll", func(t *testing.T) {

		expected := []model.Event{
			{
				ID: 0,
				Info: model.EventInfo{
					StartDate:   time.Date(2023, time.April, 15, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2023, time.June, 15, 23, 59, 59, 0, time.UTC),
					Description: "Sergey's Birthday",
				},
			},
			{
				ID: 1,
				Info: model.EventInfo{
					StartDate:   time.Date(2023, time.November, 2, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2023, time.November, 2, 23, 59, 59, 0, time.UTC),
					Description: "Tony's Birthday",
				},
			},
			{
				ID: 2,
				Info: model.EventInfo{
					StartDate:   time.Date(2023, time.June, 13, 0, 0, 0, 0, time.UTC),
					EndDate:     time.Date(2023, time.June, 13, 23, 59, 59, 0, time.UTC),
					Description: "Dasha's Birthday",
				},
			},
		}

		events, err := storage.GetAllEvents(ctx)

		require.NoError(t, err)
		require.ElementsMatch(t, expected, events)
	})

	t.Run("OK_GetByID", func(t *testing.T) {

		id := int64(1)
		expected := model.Event{
			ID: 1,
			Info: model.EventInfo{
				StartDate:   time.Date(2023, time.November, 2, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2023, time.November, 2, 23, 59, 59, 0, time.UTC),
				Description: "Tony's Birthday",
			},
		}

		event, err := storage.GetByID(ctx, id)

		require.NoError(t, err)
		require.Equal(t, expected, event)
	})

	t.Run("BAD_GetByID", func(t *testing.T) {
		id := int64(5)

		event, err := storage.GetByID(ctx, id)

		require.EqualError(t, model.ErrEventNotFound, err.Error())
		require.Equal(t, model.Event{}, event)
	})

	t.Run("OK_Remove", func(t *testing.T) {

		id := int64(2)

		err := storage.Remove(ctx, id)

		require.NoError(t, err)
	})

	t.Run("BAD_Remove", func(t *testing.T) {

		id := int64(2)

		err := storage.Remove(ctx, id)

		require.EqualError(t, model.ErrEventNotFound, err.Error())
	})

	t.Run("GetFromTo", func(t *testing.T) {
		testTable := []struct {
			name         string
			fromDate     time.Time
			toDate       time.Time
			expectedData []model.Event
		}{
			{
				name:         "case_1(empty)",
				fromDate:     time.Date(2022, time.November, 2, 0, 0, 0, 0, time.UTC),
				toDate:       time.Date(2022, time.December, 2, 0, 0, 0, 0, time.UTC),
				expectedData: []model.Event{},
			},
			{
				name:     "case_2(one_event)",
				fromDate: time.Date(2022, time.November, 2, 0, 0, 0, 0, time.UTC),
				toDate:   time.Date(2023, time.April, 17, 0, 0, 0, 0, time.UTC),
				expectedData: []model.Event{
					{
						ID: 0,
						Info: model.EventInfo{
							StartDate:   time.Date(2023, time.April, 15, 0, 0, 0, 0, time.UTC),
							EndDate:     time.Date(2023, time.June, 15, 23, 59, 59, 0, time.UTC),
							Description: "Sergey's Birthday",
						},
					},
				},
			},
			{
				name:     "case_3(everything)",
				fromDate: time.Date(2022, time.November, 2, 0, 0, 0, 0, time.UTC),
				toDate:   time.Date(2024, time.April, 17, 0, 0, 0, 0, time.UTC),
				expectedData: []model.Event{
					{
						ID: 0,
						Info: model.EventInfo{
							StartDate:   time.Date(2023, time.April, 15, 0, 0, 0, 0, time.UTC),
							EndDate:     time.Date(2023, time.June, 15, 23, 59, 59, 0, time.UTC),
							Description: "Sergey's Birthday",
						},
					},
					{
						ID: 1,
						Info: model.EventInfo{
							StartDate:   time.Date(2023, time.November, 2, 0, 0, 0, 0, time.UTC),
							EndDate:     time.Date(2023, time.November, 2, 23, 59, 59, 0, time.UTC),
							Description: "Tony's Birthday",
						},
					},
				},
			},
		}

		for _, test := range testTable {
			t.Run(test.name, func(t *testing.T) {
				events, err := storage.GetFromToEvents(ctx, test.fromDate, test.toDate)

				require.NoError(t, err)
				require.ElementsMatch(t, test.expectedData, events)
			})
		}
	})
}
