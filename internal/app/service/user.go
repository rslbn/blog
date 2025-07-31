package service

import (
	"context"
	"fmt"

	customError "github.com/rslbn/blog/internal/errors"
	"github.com/rslbn/blog/internal/model"
	db "github.com/rslbn/blog/postgres"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUserByUsername(ctx context.Context, username string) (*model.UserResponse, error)
	GetAllUsers(ctx context.Context) ([]model.UserResponse, error)
	Register(ctx context.Context, request *model.RegisterRequest) (*model.UserResponse, error)
	FindUserByUsername(ctx context.Context, username string) (*db.User, error)
}

type userService struct {
	queries *db.Queries
}

func NewUserService(queries *db.Queries) UserService {
	return &userService{queries: queries}
}

func (us *userService) GetAllUsers(ctx context.Context) ([]model.UserResponse, error) {
	users, err := us.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return []model.UserResponse{}, nil
	}
	responses := make([]model.UserResponse, 0, len(users))
	for _, user := range users {
		userResponse := &model.UserResponse{
			UserID:   uint32(user.UserID),
			Username: user.Username,
			Email:    user.Email,
		}
		responses = append(responses, *userResponse)
	}
	return responses, nil
}

func (us *userService) GetUserByUsername(ctx context.Context, username string) (*model.UserResponse, error) {
	user, err := us.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	userResponse := &model.UserResponse{
		UserID:   uint32(user.UserID),
		Username: user.Username,
		Email:    user.Email,
	}
	return userResponse, nil
}

func (us *userService) GetUserByEmail(ctx context.Context, email string) (*model.UserResponse, error) {
	user, err := us.queries.GetUserByEmail(ctx, email)
	if err != nil {
		message := fmt.Sprintf("User not found with email, %v", email)
		return nil, &customError.NotFoundError{
			Message: message,
		}
	}
	return &model.UserResponse{
		UserID:   uint32(user.UserID),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (us *userService) Register(ctx context.Context, request *model.RegisterRequest) (*model.UserResponse, error) {
	// check whether username and email is already exists
	// check username is already exist
	isExist, _ := us.queries.UserExistByUsername(ctx, request.Username)
	if isExist {
		return nil, &customError.AlreadyExistError{
			Message: "username already exist!",
		}
	}
	// check email is already exist
	isExist, _ = us.queries.UserExistsByEmail(ctx, request.Email)
	if isExist {
		return nil, &customError.AlreadyExistError{
			Message: "email already exist!",
		}
	}
	encryptPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), 1)

	if err != nil {
		return nil, err
	}
	insertUserParams := db.InsertUserParams{
		Username: request.Username,
		Password: string(encryptPass),
		Email:    request.Email,
	}
	// check password and confirmation password match
	// encrypt the password
	// save to the database
	savedUser, err := us.queries.InsertUser(ctx, insertUserParams)
	if err != nil {
		return nil, err
	}

	return &model.UserResponse{
		UserID:   uint32(savedUser.UserID),
		Username: savedUser.Username,
		Email:    savedUser.Email,
	}, nil
}

func (us *userService) FindUserByUsername(ctx context.Context, username string) (*db.User, error) {
	user, err := us.queries.GetUserByUsername(ctx, username)
	if err != nil {
		message := fmt.Sprintf("User not found with username, %v", username)
		return nil, &customError.NotFoundError{
			Message: message,
		}
	}
	return &user, nil
}
