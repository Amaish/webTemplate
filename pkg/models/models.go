package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching records found")

type Blog struct {
	ID      int
	Title   string
	Author  string
	Content string
	Created time.Time
	Expires time.Time
}
