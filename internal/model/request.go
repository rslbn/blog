package model

// User related requests
type RegisterRequest struct {
	Username             string `json:"username"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	ConfirmationPassword string `json:"confirmationPassword"`
}

// Post related request
type CreatePostRequest struct{}

type UpdatePostRequest struct{}

// Comment related request
type CreateCommentRequest struct{}
