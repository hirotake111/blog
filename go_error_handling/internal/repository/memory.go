package repository

import (
	"errors"

	"go_error_handling/internal/domain"
)

type InMemoryUserRepository struct {
	store map[string]*domain.UserInDB
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{store: make(map[string]*domain.UserInDB)}
}

func (r *InMemoryUserRepository) Add(user *domain.UserInDB) error {
	if _, ok := r.store[user.Name]; ok {
		return errors.New("user already exists")
	}
	r.store[user.Email] = user
	return nil
}

func (r *InMemoryUserRepository) Get(email string) (*domain.UserInDB, error) {
	u, ok := r.store[email]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (r *InMemoryUserRepository) Update(user *domain.UserInDB) error {
	r.store[user.Email] = user
	return nil
}
