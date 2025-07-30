package handler

import (
	"net/http"

	"github.com/rslbn/blog/internal/app/service"
	"github.com/rslbn/blog/internal/app/web"
	"github.com/rslbn/blog/internal/model"
	"github.com/rslbn/blog/internal/util"
	"github.com/rslbn/blog/internal/util/validators"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(u service.UserService) *userHandler {
	return &userHandler{u}
}

// TODO: Use github.com/go-playground/validator/v10

func (u *userHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	users, err := u.userService.GetAllUsers(r.Context())
	if err != nil {
		return err
	}
	return web.JSONResponse(w, http.StatusOK, users)
}

// TODO: choose between only get by username or both username and email
// TODO: make GetUserByUsername for userService return error
func (u *userHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) error {
	username := r.URL.Path[len("/users/"):]
	userResponse, err := u.userService.GetUserByUsername(r.Context(), username)
	if err != nil {
		return err
	}
	return web.JSONResponse(w, http.StatusOK, userResponse)
}

func (u *userHandler) Register(w http.ResponseWriter, r *http.Request) error {
	requestBody := &model.RegisterRequest{}

	err := util.DecodeJSON(r.Body, requestBody)

	if err != nil {
		return err
	}

	err = validators.ValidateRegisterRequest(requestBody)

	if err != nil {
		return err
	}

	userResponse, err := u.userService.Register(r.Context(), requestBody)
	if err != nil {
		return err
	}
	return web.JSONResponse(w, http.StatusCreated, userResponse)
}
