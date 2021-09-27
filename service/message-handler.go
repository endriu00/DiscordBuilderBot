package service

import (
	"github.com/bwmarrin/discordgo"
)

// `messageHandler` handles the event of the message creation.
// In particular, it creates a category with every channel specified
// in the configuration file of the bot.
func (bot *_bot) MessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	content := message.Content

	// If the message was sent to a channel different from the one used for categories creation
	// or if the author of the message was the bot itself
	if message.ChannelID != bot.buildChannelID || message.Author.ID == session.State.User.ID {
		return
	}

	// TODO: Check whether the message content is a topic that sticks to computer science related topics.

	// Get the list of every channel in the guild
	existingChans, err := session.GuildChannels(bot.guildID)
	if err != nil {
		bot.log.WithError(err).Error("Could not get every channel.")
		return
	}

	// Check whether the category already exists.
	// If it exists, do not create it.

	// TODO: do not stress the bot for every message with a scan of every existing channel.
	// But: see if it is convenient to sort channel names, then use a mergesort to find this.
	for _, elem := range existingChans {
		// Skip channels that are not categories
		if elem.Type != discordgo.ChannelTypeGuildCategory {
			continue
		}
		if content == elem.Name {
			bot.log.Warn("Category already exists!")
			return
		}
	}

	// Create the category first
	category, err := session.GuildChannelCreate(bot.guildID, content, discordgo.ChannelTypeGuildCategory)
	if err != nil {
		bot.log.WithError(err).WithField("userID", message.Author).Error("Could not create category.")
		return
	}

	// Create the channels according to the channel the bot has to create
	for chanName, chanType := range bot.channelsToCreate {
		_, err = session.GuildChannelCreateComplex(bot.guildID, discordgo.GuildChannelCreateData{
			Name:     chanName,
			Type:     chanType,
			ParentID: category.ID,
		})
		if err != nil {
			bot.log.WithError(err).WithField("userID", message.Author).WithField("channelName", chanName).Error("Could not create the channel.")
			return
		}
	}
}
