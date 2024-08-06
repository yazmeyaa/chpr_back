package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	config := GetConfig("./config.example.yaml")
	if config.JWT.Key != "test_key" {
		t.Errorf("wrong value in config. expected: %s, got %s", "test_key", config.JWT.Key)
	}
	fmt.Printf("String is: (%s)", config.JWT.Key)
}
