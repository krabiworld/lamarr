package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/krabiworld/lamarr/internal/bot/commands/information"
	"github.com/krabiworld/lamarr/internal/bot/commands/utilities"
	"github.com/krabiworld/lamarr/internal/bot/handlers"
	"github.com/krabiworld/lamarr/internal/bot/handlers/command"
	"github.com/krabiworld/lamarr/internal/config"
	"github.com/krabiworld/lamarr/internal/services"
	"github.com/krabiworld/lamarr/internal/types"
	"github.com/rs/zerolog/log"
)

func Start(guildService services.GuildService) {
	ownerId := config.Get().DiscordOwnerID
	guildId := config.Get().DiscordGuildID

	commandHandler := command.NewHandler(InitCommands(), InitCategories(), guildService, ownerId)
	guildEvents := handlers.NewGuildEvents(guildService)

	session, err := discordgo.New("Bot " + config.Get().DiscordToken)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	session.StateEnabled = true
	session.State.MaxMessageCount = 1000

	session.AddHandler(commandHandler.OnInteractionCreate)
	session.AddHandler(guildEvents.OnGuildCreate)

	if err = session.Open(); err != nil {
		log.Error().Err(err).Send()
		return
	}

	RegisterCommands(session, commandHandler, guildId)

	log.Info().Msg("Bot started")

	select {}
}

func InitCommands() []command.Command {
	return []command.Command{
		information.NewHelpCommand(),
		information.NewServerCommand(),
		information.NewStatsCommand(),
		information.NewUserCommand(),
		utilities.NewAvatarCommand(),
		utilities.NewRandomCommand(),
	}
}

func InitCategories() []types.Category {
	return []types.Category{
		types.CategoryInformation,
		types.CategoryUtilities,
	}
}

func RegisterCommands(session *discordgo.Session, handler *command.Handler, guildId string) {
	log.Info().Int("count", len(handler.CommandsList)).Msg("Registering commands...")

	commands := make([]*discordgo.ApplicationCommand, len(handler.CommandsList))

	for i, cmd := range handler.CommandsList {
		commands[i] = cmd.ApplicationCommand
	}

	_, err := session.ApplicationCommandBulkOverwrite(session.State.User.ID, guildId, commands)
	if err != nil {
		log.Error().Err(err).Send()
	}
}
