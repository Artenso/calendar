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

type Event struct {
	ID        int64
	Info      EventInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
