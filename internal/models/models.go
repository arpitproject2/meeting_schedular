package models

import "time"

type Meeting struct {
	ID        int
	Title     string
	RoomID    int
	StartTime time.Time
	EndTime   time.Time
	CreatedBy int
}
