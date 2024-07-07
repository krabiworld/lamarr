package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"module-go/internal/bot/commands/information"
	"module-go/internal/bot/commands/utilities"
	"module-go/internal/bot/handlers"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/cfg"
	"module-go/internal/services"
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

	commandHandler := command.NewHandler(InitCommands(), guildService, cfg.Get().DiscordOwnerID)

	session.AddHandler(guildEvents.OnGuildCreate)
	session.AddHandler(commandHandler.OnInteractionCreate)

	if err := session.Open(); err != nil {
		log.Fatal().Err(err).Send()
	}

	RegisterCommands(session, commandHandler, cfg.Get().DiscordGuildID)

	log.Info().Msg("Bot started")

	select {}
}

func InitCommands() []*command.Command {
	return []*command.Command{
		information.NewServerCommand(),
		information.NewUserCommand(),
		utilities.NewAvatarCommand(),
		utilities.NewRandomCommand(),
	}
}

func RegisterCommands(session *discordgo.Session, handler *command.Handler, guildId string) {
	for _, cmd := range handler.Commands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, guildId, cmd.ApplicationCommand)
		if err != nil {
			log.Error().Err(err).Send()
		}
	}
}
