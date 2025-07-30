package model

type UserResponse struct {
	UserID   uint32 `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
