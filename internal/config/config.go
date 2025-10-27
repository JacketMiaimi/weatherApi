package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Configs struct {
	Env         string `yaml:"env" env-default:"development"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Adress string `yaml:"address"`
	TimeOut time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"` 
}

var cfg Configs

func LoadConfig(path string) *Configs {
	op := "internal/config/config.go"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	configPath := os.Getenv("PATH_CONFIG")

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("Error opening config file: %s\n%s", err, op)
	}

	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Error reading config: %s\n%s", err, op)
	}

	return &cfg
}