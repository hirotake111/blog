package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_error_handling/internal/domain"
	"go_error_handling/internal/util"
	"net/http"
)

type UserService interface {
	SignUp(user *domain.UserSiginingUp) (*domain.UserLoggedIn, error)
}

func HandleRoot() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			encode(w, http.StatusOK, "OK")
			return
		},
	)
}

func HandleSignUp(authSvc UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := new(domain.UserSiginingUp)
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			encode(w, http.StatusBadRequest, "bad request")
			return
		}
		li, err := authSvc.SignUp(user)
		if err != nil {
			if errors.Is(err, &util.BadRequestError) {
				encode(w, http.StatusBadRequest, "bad request")
				return
			}
			encode(w, http.StatusInternalServerError, "????")
			return
		}
		encode(w, http.StatusCreated, *li)
		return
	})
}

func encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}
