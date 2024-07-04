package information

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"module-go/internal/bot/handler"
	"module-go/internal/colors"
)

type ServerCommand struct{}

func (cmd *ServerCommand) Handle(ctx *handler.CommandContext) error {
	guild, err := ctx.Guild()
	if err != nil {
		return err
	}

	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("Information about %s", guild.Name),
		Color: colors.DEFAULT,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: guild.IconURL("128"),
		},
		Image: &discordgo.MessageEmbedImage{
			URL: guild.BannerURL("512"),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("ID: %s", guild.ID),
		},
		Fields: []*discordgo.MessageEmbedField{
			cmd.MembersField(guild),
		},
	}

	return ctx.ReplyEmbed(embed)
}

func (cmd *ServerCommand) MembersField(guild *discordgo.Guild) *discordgo.MessageEmbedField {
	botCount, memberCount := 0, 0

	for _, member := range guild.Members {
		if member.User.Bot {
			botCount++
		} else {
			memberCount++
		}
	}

	return &discordgo.MessageEmbedField{
		Name:   fmt.Sprintf("Members (%d)", guild.MemberCount),
		Value:  fmt.Sprintf("Members: **%d**\nBots: **%d**", memberCount, botCount),
		Inline: true,
	}
}
