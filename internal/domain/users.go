package domain

import "time"

type User struct {
	Id        int       `json:"id,omitempty"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	RoleId    int       `json:"role_id"`
	GymId     int       `json:"gym_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
