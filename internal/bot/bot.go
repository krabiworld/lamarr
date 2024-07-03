package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/bot/commands/information"
	"module-go/internal/bot/handler"
	"module-go/internal/cfg"
)

func Start() {
	session, err := discordgo.New("Bot " + cfg.Get().DiscordToken)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	session.Identify.Intents = discordgo.IntentsAll
	session.State.MaxMessageCount = 1000

	commands := InitCommands()

	commandHandler := handler.NewCommandHandler(
		cfg.Get().DiscordOwnerID,
		cfg.Get().DiscordGuildID,
		commands,
	)
	session.AddHandler(commandHandler.OnInteraction)

	if err := session.Open(); err != nil {
		log.Fatal().Err(err).Send()
	}

	RegisterCommands(session, commandHandler)

	log.Info().Msg("Bot started")

	select {}
}

func InitCommands() []*handler.Command {
	return []*handler.Command{
		{
			Command: &discordgo.ApplicationCommand{
				Name:        "server",
				Description: "Show information about current server.",
			},
			Category:          handler.INFORMATION,
			OwnerCommand:      false,
			ModerationCommand: false,
			Hidden:            false,
			Handler:           &information.ServerCommand{},
		},
	}
}

func RegisterCommands(session *discordgo.Session, handler *handler.CommandHandler) {
	for _, command := range handler.Commands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, handler.ForceGuildID, command.Command)
		if err != nil {
			log.Error().Err(err).Send()
		}
	}
}
