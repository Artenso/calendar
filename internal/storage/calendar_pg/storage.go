package calendar_pg

import (
	"context"
	"time"

	"github.com/Artenso/calendar/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

const (
	table = "calendar"

	idCol          = "id"
	descriptionCol = "description"
	startDateCol   = "startDate"
	endDateCol     = "endDate"
	createdAtCol   = "createdAt"
	updatedAtCol   = "updatedAt"
)

type storage struct {
	dbConn *pgx.Conn
}

func New(dbConn *pgx.Conn) *storage {
	return &storage{
		dbConn: dbConn,
	}
}

func (s *storage) Add(ctx context.Context, info *model.EventInfo) (*model.Event, error) {
	builder := sq.Insert(table).
		Columns(descriptionCol, startDateCol, endDateCol, createdAtCol).
		Values(info.Description, info.StartDate, info.EndDate, time.Now().UTC()).
		Suffix("RETURNING id, createdAt").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build sql query")
	}

	event := &model.Event{
		Info: *info,
	}

	err = s.dbConn.QueryRow(ctx, query, args...).Scan(&event.ID, &event.CreatedAt)
	if err != nil {
		return nil, err
	}

	return event, nil
}
func (s *storage) Remove(ctx context.Context, eventID int64) error {
	return nil
}
func (s *storage) Edit(ctx context.Context, eventID int64, info *model.UpdateEventInfo) (*model.Event, error) {
	return nil, nil
}
func (s *storage) GetAllEvents(ctx context.Context) ([]*model.Event, error) {
	return nil, nil
}
func (s *storage) GetFromToEvents(ctx context.Context, from time.Time, to time.Time) ([]*model.Event, error) {
	return nil, nil
}
func (s *storage) GetByID(ctx context.Context, eventID int64) (*model.Event, error) {
	return nil, nil
}
