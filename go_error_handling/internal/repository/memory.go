package repository

import (
	"errors"
	"log"

	"go_error_handling/internal/domain"
)

type InMemoryUserRepository struct {
	store map[string]*domain.UserWithPasswordHash
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{store: make(map[string]*domain.UserWithPasswordHash)}
}

func (r *InMemoryUserRepository) Add(user *domain.UserWithPasswordHash) error {
	if _, ok := r.store[user.Name]; ok {
		return errors.New("user already exists")
	}
	r.store[user.Email] = user
	for e, u := range r.store {
		log.Printf("%s - %s\n", e, u.Name)
	}
	log.Printf("length: %d\n", len(r.store))
	return nil
}

func (r *InMemoryUserRepository) Get(id string) (*domain.UserWithPasswordHash, error) {
	u, ok := r.store[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}
