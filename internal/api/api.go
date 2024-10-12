package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/krabiworld/lamarr/internal/cfg"
	"github.com/rs/zerolog/log"
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

	listenConfig := fiber.ListenConfig{
		DisableStartupMessage: true,
	}

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Lamarr API")
	})

	app.Get("/commands", func(c fiber.Ctx) error {
		return nil
	})
	app.Get("/stats", func(c fiber.Ctx) error {
		return nil
	})

	go func() {
		log.Fatal().Err(app.Listen(cfg.Get().ApiAddr, listenConfig)).Send()
	}()

	log.Info().Msg("API started")

	select {}
}
