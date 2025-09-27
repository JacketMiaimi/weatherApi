package config

import (
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type Configs struct {
	Env         string `yaml:"env" env-default:"development"`
	StoragePath string `yaml:"storagePath" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Adress string `yaml:"address"`
}

var cfg Configs

func LoadConfig(path string) *Configs {
	// op := "internal/config/config.go"

	if _, err := os.Stat(path); err != nil {
		log.Fatalf("Error opening config file: %s", err)
	}

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	return &cfg
}