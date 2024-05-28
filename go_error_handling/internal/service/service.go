package service

import (
	"errors"
	"go_error_handling/internal/domain"

	// "go_error_handling/internal/util"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Get(email string) (*domain.UserInDB, error)
	Add(user *domain.UserInDB) error
	Update(user *domain.UserInDB) error
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) SignUp(user *domain.UserSigningUp) (*domain.UserLoggedIn, error) {
	if u, _ := s.repository.Get(user.Email); u != nil {
		return nil, errors.New("user already exists")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return nil, err
	}
	udb := user.ToUserInDB(hash, time.Now())
	if err = s.repository.Add(udb); err != nil {
		return nil, err
	}
	return udb.ToLoggedIn(), nil
}

func (s *UserService) Login(user *domain.UserLoggingIn) (*domain.UserLoggedIn, error) {
	udb, err := s.repository.Get(user.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(udb.PasswordHash, []byte(user.Password))
	if err != nil {
		return nil, err
	}
	udb.LoggedInAt = time.Now()
	err = s.repository.Update(udb)
	if err != nil {
		return nil, err
	}
	loggedIn := udb.ToLoggedIn()
	return loggedIn, nil

}
