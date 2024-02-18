package model

import "time"

type EventInfo struct {
	StartDate   time.Time
	EndDate     time.Time
	Description string
}

type Event struct {
	ID   int64
	Info EventInfo
}
