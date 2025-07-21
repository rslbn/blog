package service

import (
	"context"

	"github.com/rslbn/blog/internal/model"
)

type UserService interface {
	GetUserByUsername(ctx context.Context, username string) *model.UserResponse
	Login(ctx context.Context, request *model.LoginRequest) *model.LoginResponse
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (us *userService) GetUserByUsername(ctx context.Context, username string) *model.UserResponse {
	return &model.UserResponse{}
}

func (us *userService) Login(ctx context.Context, request *model.LoginRequest) *model.LoginResponse {
	return &model.LoginResponse{}
}
