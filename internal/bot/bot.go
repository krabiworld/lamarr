package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/bot/commands/information"
	"module-go/internal/cfg"
)

func Start() {
	session, err := discordgo.New("Bot " + cfg.Get().DiscordToken)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	session.AddHandler(information.Server)

	if err := session.Open(); err != nil {
		log.Fatal().Err(err).Send()
	}

	log.Info().Msg("Bot started")

	select {}
}
