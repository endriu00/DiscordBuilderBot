package service

import (
	"errors"
	dgo "github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

var ErrNotCommand = errors.New("message content is not a command")

// `Config` is the configuration for the bot.
type Config struct {
	GuildID          string
	BuildChannelID   string
	Log              *logrus.Entry
	ChannelsToCreate map[string]dgo.ChannelType
}

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

func New(cfg Config) *_bot {
	return &_bot{
		guildID:          cfg.GuildID,
		buildChannelID:   cfg.BuildChannelID,
		log:              cfg.Log,
		channelsToCreate: cfg.ChannelsToCreate,
	}
}
