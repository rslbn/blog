package handler

import (
	"net/http"

	"github.com/rslbn/blog/internal/app/service"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(a service.AuthService) *authHandler {
	return &authHandler{a}
}

func (a *authHandler) Login(w http.ResponseWriter, r *http.Request) {

}
