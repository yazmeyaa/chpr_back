package authentication

import (
	"github.com/yazmeyaa/chat_backend/internal/data/request"
	"github.com/yazmeyaa/chat_backend/internal/services/jwt"
)

type AuthenticationService struct {
	jwt *jwt.JWTService
}

func New(jwt *jwt.JWTService) *AuthenticationService {
	return &AuthenticationService{
		jwt,
	}
}

func (a AuthenticationService) Login(user request.LoginRequest) (token string) {
	options := jwt.CreateTokenOptions{
		Name: user.Username,
	}
	token = a.jwt.GenerateToken(options)

	return token
}
