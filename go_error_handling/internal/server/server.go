package server

import (
	"net"
	"net/http"

	"go_error_handling/internal/config"
	"go_error_handling/internal/handler"
	"go_error_handling/internal/service"
)

type Server struct {
	addr    string
	handler http.Handler
}

func NewServer(config *config.Config, as *service.UserService) *http.Server {
	addr := net.JoinHostPort(config.Host, config.Port)
	mux := http.NewServeMux()
	addRoute(mux, config, as)

	var handler http.Handler = mux
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}

func addRoute(mux *http.ServeMux, _ *config.Config, as *service.UserService) {
	mux.Handle("GET /", handler.HandleRoot())
	mux.Handle("POST /signup", handler.HandleSignUp(as))
	mux.Handle("POST /login", handler.HandleLogin(as))

}
