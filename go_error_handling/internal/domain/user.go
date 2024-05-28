package domain

import (
	"errors"
	"time"
)

const (
	minEmailLength    = 6
	minPasswordLength = 4
	minNameLength     = 4
)

type UserInDB struct {
	Email        string
	Name         string
	PasswordHash []byte
	LoggedInAt   time.Time
}

func (u *UserInDB) ToLoggedIn() *UserLoggedIn {
	return &UserLoggedIn{
		Email:      u.Email,
		Name:       u.Name,
		LoggedInAt: u.LoggedInAt,
	}
}

type UserLoggingIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserLoggingIn(email, password string) (*UserLoggingIn, error) {
	if len(email) < minEmailLength || len(password) < minPasswordLength {
		return nil, errors.New("either email, or password is too short")
	}
	return &UserLoggingIn{Email: email, Password: password}, nil
}

type UserLoggedIn struct {
	Email      string    `json:"string"`
	Name       string    `json:"name"`
	LoggedInAt time.Time `json:"logged_in_at"`
}

type UserSigningUp struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *UserSigningUp) ToUserInDB(hash []byte, loggedInAt time.Time) *UserInDB {
	return &UserInDB{
		Email:        u.Email,
		Name:         u.Name,
		PasswordHash: hash,
		LoggedInAt:   loggedInAt,
	}
}

func NewUserSigningUp(name string, email string, password string) (*UserSigningUp, error) {
	if len(name) < 4 || len(email) < 8 || len(password) < 4 {
		return nil, errors.New("either name, email, or password is too short")
	}
	return &UserSigningUp{
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
