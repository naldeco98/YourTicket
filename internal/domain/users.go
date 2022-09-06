package domain

import "time"

type User struct {
	Id        int
	Username  string
	Password  string
	RoleId    int
	GymId     int
	CreatedAt time.Time
}
