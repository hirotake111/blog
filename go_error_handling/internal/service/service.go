package service

import (
	"go_error_handling/internal/domain"
	"go_error_handling/internal/util"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Get(email string) (*domain.UserWithPasswordHash, error)
	Add(u *domain.UserWithPasswordHash) error
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) SignUp(user *domain.UserSiginingUp) (*domain.UserLoggedIn, error) {
	_, err := s.repository.Get(user.Email)
	if err == nil {
		return nil, &util.BadRequestError
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return nil, &util.InternalServerError
	}
	newUser := &domain.UserWithPasswordHash{
		Email:        user.Email,
		Name:         user.Name,
		PasswordHash: hash,
	}
	err = s.repository.Add(newUser)
	if err != nil {
		return nil, &util.InternalServerError
	}
	return user.ToLogIn(time.Now()), nil
}
