package service

import (
	"context"

	customError "github.com/rslbn/blog/internal/errors"
	"github.com/rslbn/blog/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(c context.Context, r *model.LoginRequest) (*model.LoginResponse, error)
}

type authService struct {
	userService UserService
	jwtService  JwtService
}

func NewAuthService(us UserService, js JwtService) AuthService {
	return &authService{us, js}
}

func (a *authService) Login(c context.Context, r *model.LoginRequest) (*model.LoginResponse, error) {
	// find user by username
	user, err := a.userService.FindUserByUsername(c, r.Username)
	if err != nil {
		return nil, &customError.InvalidCredentialsError{
			Message: "invalid credentials",
		}
	}

	// match the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		return nil, &customError.InvalidCredentialsError{
			Message: "invalid credentials",
		}
	}

	token, err := a.jwtService.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		ID: uint32(user.UserID), Email: user.Email, Username: user.Username, Token: token,
	}, nil
}
