package cfg

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
	DiscordGuildID uint64 `env:"DISCORD_GUILD_ID"`
	DiscordBotID   uint64 `env:"DISCORD_BOT_ID"`
	DiscordToken   string `env:"DISCORD_TOKEN"`

	// Database
	DBDatabase string `env:"DB_DATABASE"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBHostname string `env:"DB_HOSTNAME"`
	DBPort     string `env:"DB_PORT"`
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
