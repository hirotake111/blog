package main

import (
	"log"

	"go_error_handling/internal/config"
	"go_error_handling/internal/repository"
	"go_error_handling/internal/server"
	"go_error_handling/internal/service"
)

func main() {
	config := config.NewConfig("127.0.0.1", "8080")
	ur := repository.NewInMemoryUserRepository()
	us := service.NewUserService(ur)
	svr := server.NewServer(config, us)
	log.Printf("Listening on %s\n", svr.Addr)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
