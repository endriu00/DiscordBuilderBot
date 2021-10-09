package handler

import (
	"errors"
	dgo "github.com/bwmarrin/discordgo"
	"github.com/endriu00/DiscordBuilderBot/service/db"
	"github.com/sirupsen/logrus"
)

// Constants for message strings.
const (
	// `doneMessage` is the message indicating the channel was successfully created.
	doneMessage = "Already done, pal!"

	// `errorMessage` is the message indicating an error has occurred.
	errorMessage = "Looks like I am broken."

	// `notCommandMessage` is the message indicating an error made by the user.
	notCommandMessage = "Names not starting with ! are not considered."

	// `catExistingMessage` is the message indicating the category already exists.
	catExistsMessage = "This category already exists!"

	// `messageSentPoints` represents the points the user receives when sending a message.
	messageSentPoints = 1

	// `messageReplyPoints` represents the points the user receives when replying
	// to a message.
	messageReplyPoints = 2
)

// `ErrNotCommand` indicated the message content is not a command meant for the bot.
var ErrNotCommand = errors.New("message content is not a command")

// `Config` is the configuration for the bot.
type Config struct {
	GuildID          string
	BuildChannelID   string
	Log              *logrus.Entry
	ChannelsToCreate map[string]dgo.ChannelType
	DB               *db.DB
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

	// `db` is the db the bot interacts to.
	db *db.DB
}

// `New` creates a `_bot` from a configuration `cfg`.
func New(cfg Config) *_bot {
	// TODO check whether everything has been initialized
	return &_bot{
		guildID:          cfg.GuildID,
		buildChannelID:   cfg.BuildChannelID,
		log:              cfg.Log,
		channelsToCreate: cfg.ChannelsToCreate,
		db:               cfg.DB,
	}
}
