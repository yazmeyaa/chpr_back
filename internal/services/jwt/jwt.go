package jwt

import (
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yazmeyaa/chat_backend/internal/services/config"
)

type JWTService struct {
	config config.Config
}

func New(config *config.Config) *JWTService {
	return &JWTService{
		config: *config,
	}
}

type CreateTokenOptions struct {
	Name string
}

func (j JWTService) GenerateToken(options CreateTokenOptions) string {
	claims := jwt.MapClaims{
		"sub": options.Name,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iss": "chat",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString([]byte(j.config.JWT.Key))
	if err != nil {
		slog.Error(err.Error())
	}

	return str
}
