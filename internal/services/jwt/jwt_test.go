package jwt

import (
	"testing"

	"github.com/yazmeyaa/chat_backend/internal/services/config"
)

func TestJWT(t *testing.T) {
	config := config.New()
	jwtService := New(config)

	token := jwtService.GenerateToken(CreateTokenOptions{
		Name: "yazmeyaa",
	})

	if token == "" {
		t.Error("Recieved empty token.")
	}

}
