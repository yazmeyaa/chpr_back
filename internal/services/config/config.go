package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	JWT struct {
		Key string `yaml:"key"`
	} `yaml:"jwt"`
}

func GetConfig(path string) *Config {
	config := Config{}

	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("Wrong file path provided (%s)", path)
	}
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatalf("Cannot open config file %s", absPath)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Cannot decode config file")
	}
	return &config
}
