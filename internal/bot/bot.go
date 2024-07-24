package bot

import (
	"context"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/snowflake/v2"
	"github.com/rs/zerolog/log"
	"module-go/internal/bot/commands/information"
	"module-go/internal/bot/commands/utilities"
	"module-go/internal/bot/handlers"
	"module-go/internal/bot/handlers/command"
	"module-go/internal/cfg"
	"module-go/internal/services"
	"module-go/internal/types"
	"time"
)

func Start(guildService services.GuildService) {
	ownerId := snowflake.MustParse(cfg.Get().DiscordOwnerID)
	guildId := snowflake.MustParse(cfg.Get().DiscordGuildID)

	commandHandler := command.NewHandler(InitCommands(), InitCategories(), guildService, ownerId)
	guildEvents := handlers.NewGuildEvents(guildService)

	client, err := disgo.New(
		cfg.Get().DiscordToken,
		bot.WithGatewayConfigOpts(gateway.WithIntents(gateway.IntentsAll)),
		bot.WithCacheConfigOpts(cache.WithCaches(cache.FlagGuilds, cache.FlagMembers, cache.FlagPresences)),
		bot.WithEventListenerFunc(commandHandler.OnInteractionCreate),
		bot.WithEventListenerFunc(guildEvents.OnGuildCreate),
	)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.OpenGateway(ctx); err != nil {
		log.Fatal().Err(err).Send()
	}

	RegisterCommands(client, commandHandler, guildId)

	log.Info().Msg("Bot started")

	select {}
}

func InitCommands() []command.Command {
	return []command.Command{
		information.NewHelpCommand(),
		information.NewServerCommand(),
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

func RegisterCommands(client bot.Client, handler *command.Handler, guildId snowflake.ID) {
	log.Info().Int("count", len(handler.CommandsList)).Msg("Registering commands...")

	commands := make([]discord.ApplicationCommandCreate, len(handler.CommandsList))

	for i, cmd := range handler.CommandsList {
		commands[i] = cmd.ApplicationCommand
	}

	_, err := client.Rest().SetGuildCommands(client.ApplicationID(), guildId, commands)
	if err != nil {
		log.Error().Err(err).Send()
	}
}
