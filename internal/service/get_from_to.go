package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Artenso/calendar/internal/model"
)

func (s *service) GetFromToEvents(ctx context.Context, fromDate, toDate time.Time) ([]model.Event, error) {
	if fromDate.After(toDate) {
		return nil, fmt.Errorf("wrong from/to format, 'to' must be later than 'from'")
	}
	utcFrom := fromDate.UTC()
	utcTo := toDate.UTC()
	events := s.eventsStorage.GetFromToEvents(ctx, utcFrom, utcTo)

	return events, nil
}
