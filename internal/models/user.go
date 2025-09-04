package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Picture  string `json:"picture"`
	Bio      string `json:"bio"`
	Online   bool   `json:"online"`
}

var TestUsers = []User{
	{
		ID:       "1",
		Username: "callum",
		Email:    "thisisanemail@gmail.com",
		Password: "Password123",
		Picture:  "",
		Bio:      "i love playrates!!!!",
		Online:   true,
	},
	{
		ID:       "2",
		Username: "testaccount",
		Email:    "otheremail@gmail.com",
		Password: "Password321",
		Picture:  "",
		Bio:      "",
		Online:   false,
	},
}
