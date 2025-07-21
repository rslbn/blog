package app

import (
	"net/http"

	"github.com/rslbn/blog/internal/app/handler"
	"github.com/rslbn/blog/internal/app/service"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService(userService)
	authHandler := handler.NewAuthHandler(authService)

	mux.HandleFunc("GET /users/{username}", userHandler.GetUserByUsername)
	mux.HandleFunc("POST /auth/login", authHandler.Login)
	return mux
}
