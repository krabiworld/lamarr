package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/bot/commands/information"
	"module-go/internal/bot/handlers"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/cfg"
	"module-go/internal/services"
	"module-go/internal/types"
)

func Start(guildService services.GuildService) {
	session, err := discordgo.New("Bot " + cfg.Get().DiscordToken)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	session.Identify.Intents = discordgo.IntentsAll
	session.StateEnabled = true
	session.State.MaxMessageCount = 100

	guildEvents := handlers.NewGuildEvents(guildService)

	commandHandler := command.NewCommandHandler(InitCommands(), guildService)

	session.AddHandler(guildEvents.OnGuildCreate)
	session.AddHandler(commandHandler.OnMessage)

	if err := session.Open(); err != nil {
		log.Fatal().Err(err).Send()
	}

	log.Info().Msg("Bot started")

	select {}
}

func InitCommands() map[string]*command.Command {
	return map[string]*command.Command{
		"server": {
			Name:              "server",
			Description:       "Information about server",
			Category:          types.INFORMATION,
			OwnerCommand:      false,
			ModerationCommand: false,
			Hidden:            false,
			Arguments:         nil,
			Handler:           &information.ServerCommand{},
		},
	}
}
