package embed

import (
	"github.com/bwmarrin/discordgo"
)

type Builder struct {
	embed *discordgo.MessageEmbed
}

func New() *Builder {
	return &Builder{embed: &discordgo.MessageEmbed{}}
}

func (b *Builder) Title(title string) *Builder {
	b.embed.Title = title
	return b
}

func (b *Builder) Description(description string) *Builder {
	b.embed.Description = description
	return b
}

func (b *Builder) Color(color int) *Builder {
	b.embed.Color = color
	return b
}

func (b *Builder) Footer(footer string) *Builder {
	b.embed.Footer = &discordgo.MessageEmbedFooter{
		Text: footer,
	}
	return b
}

func (b *Builder) Image(url string) *Builder {
	b.embed.Image = &discordgo.MessageEmbedImage{
		URL: url,
	}
	return b
}

func (b *Builder) Thumbnail(url string) *Builder {
	b.embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL: url,
	}
	return b
}

func (b *Builder) Video(url string) *Builder {
	b.embed.Video = &discordgo.MessageEmbedVideo{
		URL: url,
	}
	return b
}

func (b *Builder) Provider(name, url string) *Builder {
	b.embed.Provider = &discordgo.MessageEmbedProvider{
		Name: name,
		URL:  url,
	}
	return b
}

func (b *Builder) Author(name, url string) *Builder {
	b.embed.Author = &discordgo.MessageEmbedAuthor{
		Name:    name,
		IconURL: url,
	}
	return b
}

func (b *Builder) Field(name, value string, inline bool) *Builder {
	b.embed.Fields = append(b.embed.Fields, &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})
	return b
}

func (b *Builder) RawField(field *discordgo.MessageEmbedField) *Builder {
	b.embed.Fields = append(b.embed.Fields, field)
	return b
}

func (b *Builder) Build() *discordgo.MessageEmbed {
	return b.embed
}
