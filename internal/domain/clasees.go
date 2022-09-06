package domain

import "time"

type Class struct {
	Id        int
	Type      string
	Date      time.Time
	CreatedAt time.Time
}
