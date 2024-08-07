package authentication

import (
	"testing"

	"github.com/yazmeyaa/chat_backend/internal/data/request"
	"github.com/yazmeyaa/chat_backend/internal/services/config"
	"github.com/yazmeyaa/chat_backend/internal/services/jwt"
)

func TestAuth(t *testing.T) {
	config := config.New()
	jwtService := jwt.New(config)
	authService := New(jwtService)

	userRequest := request.LoginRequest{Username: "yazmeyaa"}
	token := authService.Login(userRequest)

	if token == "" {
		t.Error("Empty token recieved.")
	}
}
