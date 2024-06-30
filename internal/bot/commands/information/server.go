package information

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func Server(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild, err := s.Guild(i.GuildID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch guild")
		return
	}

	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("Information about %s", guild.Name),
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to interact with interaction")
	}
}
