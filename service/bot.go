package service

import (
	dgo "github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// `_bot` represents the bot.
type _bot struct {
	// `guildID` is the ID of the guild (also called server).
	guildID string

	// `buildChannelID` is the ID of the channel that lets anyone create
	// a new category.
	buildChannelID string

	// `log` is the logger
	log *logrus.Entry

	// `channelsToCreate` represents the name of the channels to be created.
	channelsToCreate map[string]dgo.ChannelType
}