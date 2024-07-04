package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/bot/commands/information"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/bot/types"
	"module-go/internal/cfg"
)

func Start() {
	session, err := discordgo.New("Bot " + cfg.Get().DiscordToken)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	session.Identify.Intents = discordgo.IntentsAll
	session.StateEnabled = true
	session.State.MaxMessageCount = 100

	commands := InitCommands()
	commandHandler := command.NewCommandHandler(commands)
	session.AddHandler(commandHandler.OnInteraction)

	if err := session.Open(); err != nil {
		log.Fatal().Err(err).Send()
	}

	RegisterCommands(session, commandHandler, cfg.Get().DiscordGuildID)

	log.Info().Msg("Bot started")

	select {}
}

func InitCommands() []*command.Command {
	return []*command.Command{
		{
			Command: &discordgo.ApplicationCommand{
				Name:        "server",
				Description: "Show information about current server.",
			},
			Category:          types.INFORMATION,
			OwnerCommand:      false,
			ModerationCommand: false,
			Hidden:            false,
			Handler:           &information.ServerCommand{},
		},
	}
}

func RegisterCommands(session *discordgo.Session, commandHandler *command.Handler, guildId string) {
	for _, cmd := range commandHandler.Commands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, guildId, cmd.Command)
		if err != nil {
			log.Error().Err(err).Send()
		}
	}
}
