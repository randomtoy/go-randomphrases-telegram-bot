package config

import (
	"context"

	"github.com/heetch/confita"
)

type Database struct {
	Username string
	Password string
	Database string
	URI      string
}

type Telegram struct {
	ApiToken string
}
type Config struct {
	Database Database
	Telegram Telegram
}

func NewConfig() (*Config, error) {
	loader := confita.NewLoader()

	cfg := Config{}

	err := loader.Load(context.Background(), &cfg)

	if err != nil {
		return &Config{}, err
	}
	return &cfg, nil

}
