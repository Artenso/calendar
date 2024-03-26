package model

import (
	"database/sql"
	"time"
)

type EventInfo struct {
	StartDate   time.Time
	EndDate     time.Time
	Description string
}

type UpdateEventInfo struct {
	StartDate   sql.NullTime
	EndDate     sql.NullTime
	Description sql.NullString
}

type Event struct {
	ID        int64
	Info      EventInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
