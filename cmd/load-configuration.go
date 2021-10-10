package main

import (
	"errors"
	dgo "github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
)

// `BotConfiguration` is the bot configuration.
// Used for communicating with internal bot structure and configuration.
type BotConfiguration struct {
	GuildID            string
	BuildChannelID     string
	GetPointsChannelID string
	ChannelsToCreate   map[string]dgo.ChannelType
}

// `loadConfigurationFromEnv` load the configuration of the bot
// from the environment variables.
func loadConfigurationFromEnv() (BotConfiguration, error) {
	err := godotenv.Load()
	if err != nil {
		return BotConfiguration{}, errors.New("cannot load environment variables")
	}

	// Fetch the guildID
	guildID, _ := os.LookupEnv("GUILD_ID")
	if guildID == "" {
		return BotConfiguration{}, errors.New("environment variable does not exists or it is not set")
	}

	// Fetch the parameters for the build category channel
	buildChannelID, _ := os.LookupEnv("BUILD_CHANNEL_ID")
	if buildChannelID == "" {
		return BotConfiguration{}, errors.New("environment variable does not exists or it is not set")
	}
	channelsToCreate := make(map[string]dgo.ChannelType)
	channelsToCreate["useful-resources"] = dgo.ChannelTypeGuildText
	channelsToCreate["project-and-ideas"] = dgo.ChannelTypeGuildText
	channelsToCreate["general-discussion"] = dgo.ChannelTypeGuildText
	channelsToCreate["let's talk"] = dgo.ChannelTypeGuildVoice

	// Fetch the parameters for the get points channel
	getPointsChannelID, _ := os.LookupEnv("GET_POINTS_CHANNEL_ID")
	if getPointsChannelID == "" {
		return BotConfiguration{}, errors.New("environment variable does not exists or it is not set")
	}

	return BotConfiguration{
		GuildID:            guildID,
		BuildChannelID:     buildChannelID,
		GetPointsChannelID: getPointsChannelID,
		ChannelsToCreate:   channelsToCreate,
	}, nil

}
