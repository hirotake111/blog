package domain

import "time"

type UserWithPasswordHash struct {
	Email        string
	Name         string
	PasswordHash []byte
}

type UserLoggedIn struct {
	Email      string    `json:"string"`
	Name       string    `json:"name"`
	LoggedInAt time.Time `json:"logged_in_at"`
}

type UserSiginingUp struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *UserSiginingUp) ToLogIn(now time.Time) *UserLoggedIn {
	return &UserLoggedIn{
		Email:      u.Email,
		Name:       u.Name,
		LoggedInAt: now,
	}
}
