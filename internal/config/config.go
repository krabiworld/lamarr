package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	// Common
	AppName  string `env:"APP_NAME"`
	LogLevel string `env:"LOG_LEVEL"`
	Debug    bool   `env:"DEBUG" envDefault:"false"`

	// API
	ApiAddr string `env:"API_ADDR"`

	// Discord
	DiscordOwnerID string `env:"DISCORD_OWNER_ID"`
	DiscordGuildID string `env:"DISCORD_GUILD_ID"`
	DiscordToken   string `env:"DISCORD_TOKEN"`

	// Database
	DatabaseType string `env:"DATABASE_TYPE"`
	DatabaseDSN  string `env:"DATABASE_DSN"`
}

var cfg *Config

func Init() {
	// Load env variables from file
	_ = godotenv.Load()

	cfg = &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func Get() *Config {
	return cfg
}
