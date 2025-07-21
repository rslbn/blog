package handler

import (
	"net/http"

	"github.com/rslbn/blog/internal/app/service"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(u service.UserService) *userHandler {
	return &userHandler{u}
}

// TODO: choose between only get by username or both username and email
func (u *userHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/users/"):]
	w.Write([]byte(username))
}

func (u *userHandler) Register(w http.ResponseWriter, r *http.Request) {

}

func (u *userHandler) Login(w http.ResponseWriter, r *http.Request) {

}
