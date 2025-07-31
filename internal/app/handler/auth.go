package handler

import (
	"fmt"
	"net/http"

	"github.com/rslbn/blog/internal/app/service"
	"github.com/rslbn/blog/internal/app/web"
	"github.com/rslbn/blog/internal/model"
	"github.com/rslbn/blog/internal/util"
	"github.com/rslbn/blog/internal/util/validators"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(a service.AuthService) *authHandler {
	return &authHandler{a}
}

// TODO: Add more functionality or features that makes login safe
// TODO: Integrate with AuthService
func (a *authHandler) Login(w http.ResponseWriter, r *http.Request) error {
	requestBody := model.LoginRequest{}
	err := util.DecodeJSON(r.Body, &requestBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = validators.ValidateLoginRequest(&requestBody)

	if err != nil {
		return err
	}

	// Process the request to authService.Login
	loginResponse, err := a.authService.Login(r.Context(), &requestBody)
	if err != nil {
		return err
	}
	return web.JSONResponse(w, http.StatusOK, loginResponse)
}
