package model

import (
	"errors"
)

var (
	ErrDateBusy      = errors.New("the date is busy with another event")
	ErrEventNotFound = errors.New("event not found")
)
