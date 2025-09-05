package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"createdAt"`
}
