package model

// User related requests
type RegisterRequest struct {
	Username, Email, Password, ConfirmationPassword string
}

// Post related request
type CreatePostRequest struct{}

type UpdatePostRequest struct{}

// Comment related request
type CreateCommentRequest struct{}
