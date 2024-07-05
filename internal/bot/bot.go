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

	commands := InitCommands()
	commandHandler := command.NewCommandHandler(commands, guildService)

	session.AddHandler(guildEvents.OnGuildCreate)
	session.AddHandler(commandHandler.OnMessage)
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
			Name:              "server",
			Description:       "Information about server",
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
		applicationCommand := &discordgo.ApplicationCommand{
			Name:        cmd.Name,
			Description: cmd.Description,
		}
		_, err := session.ApplicationCommandCreate(session.State.User.ID, guildId, applicationCommand)
		if err != nil {
			log.Error().Err(err).Send()
		}
	}
}
