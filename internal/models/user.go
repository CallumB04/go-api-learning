package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"createdAt"`
}

var TestUsers = []User{
	{
		ID:        "1",
		Email:     "thisisanemail@gmail.com",
		Username:  "callum",
		Bio:       "This is a test bio",
		CreatedAt: time.Now().UTC(),
	},
	{
		ID:        "2",
		Email:     "otheremail@gmail.com",
		Username:  "testperson",
		Bio:       "",
		CreatedAt: time.Now().UTC(),
	},
}
