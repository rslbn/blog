package service

type AuthService interface {
}

type authService struct {
	userService UserService
}

func NewAuthService(us UserService) AuthService {
	return &authService{us}
}
