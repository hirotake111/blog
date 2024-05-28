package handler

import (
	"encoding/json"
	"fmt"
	"go_error_handling/internal/domain"

	// "go_error_handling/internal/util"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type UserService interface {
	SignUp(user *domain.UserSigningUp) (*domain.UserLoggedIn, error)
	Login(user *domain.UserLoggingIn) (*domain.UserLoggedIn, error)
}

func HandleRoot() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			encodeJson(w, http.StatusOK, "OK")
			return
		},
	)
}

func HandleSignUp(us UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b := new(struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		})
		if err := json.NewDecoder(r.Body).Decode(b); err != nil {
			encodeJson(w, http.StatusBadRequest, Response{
				Message: "bad request",
			})
			return
		}
		userSigningUp, err := domain.NewUserSigningUp(b.Name, b.Email, b.Password)
		if err != nil {
			encodeJson(w, http.StatusBadRequest, Response{
				Message: "bad request",
			})
			return
		}
		userLoggedIn, err := us.SignUp(userSigningUp)
		if err != nil {
			encodeJson(w, http.StatusBadRequest, Response{
				Message: "bad request",
			})
			return
		}
		if err = encodeJson(w, http.StatusCreated, Response{
			Message: fmt.Sprintf("user %s created", userLoggedIn.Name),
		}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("internal server error")
			return
		}
		return
	})
}

func HandleLogin(us UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b := new(struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		})
		if err := json.NewDecoder(r.Body).Decode(b); err != nil {
			encodeJson(w, http.StatusBadRequest, Response{Message: "bad request"})
			return
		}
		loggingIn, err := domain.NewUserLoggingIn(b.Email, b.Password)
		if err != nil {
			encodeJson(w, http.StatusBadRequest, Response{Message: "bad request"})
			return
		}
		u, err := us.Login(loggingIn)
		if err != nil {
			encodeJson(w, http.StatusBadRequest, Response{Message: "bad request"})
			return
		}
		encodeJson(w, http.StatusOK, Response{
			Message: fmt.Sprintf("hello %s", u.Email),
		})
		return
	})
}

func encodeJson[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(&v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}
