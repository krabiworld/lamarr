package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
	"module-go/internal/cfg"
)

func Start() {
	appName := cfg.Get().AppName

	app := fiber.New(fiber.Config{
		ServerHeader: appName,
		ErrorHandler: nil,
		AppName:      appName,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	app.Get("/commands", func(c fiber.Ctx) error {
		return nil
	})

	log.Fatal().Err(app.Listen(cfg.Get().ApiAddr)).Send()
}
