package handler

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// `messageHandler` handles the event of the message creation.
// In particular, it creates a category with every channel specified
// in the configuration file of the bot.
func (bot *_bot) MessageBuildCategoryHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	// If the message was sent to a channel different from the one used for categories creation
	// or if the author of the message was the bot itself
	if message.ChannelID != bot.buildChannelID || message.Author.ID == session.State.User.ID {
		return
	}
	// Check whether the content of the message has the correct starting character
	content, err := bot.SanitizeCommand(message.Content)
	if err == ErrNotCommand {
		bot.log.WithError(err).WithField("userID", message.Author.ID).Error("Not a command meant for the bot.")
		if err = bot.SendMessage(notCommandMessage, bot.buildChannelID, session); err != nil {
			bot.log.WithError(err).Error("Error sending message")
		}
		return
	}

	// TODO: Check whether the message content is a topic that sticks to computer science related topics.

	// Get the list of every channel in the guild
	existingChans, err := session.GuildChannels(bot.guildID)
	if err != nil {
		bot.log.WithError(err).Error("Could not get every channel.")
		if err = bot.SendMessage(errorMessage, bot.buildChannelID, session); err != nil {
			bot.log.WithError(err).Error("Error sending message")
		}
		return
	}

	// Check whether the category already exists.
	// If it exists, do not create it. Names are case insensitive.

	// TODO: do not stress the bot for every message with a scan of every existing channel.
	// But: see if it is convenient to sort channel names, then use a mergesort to find this.
	for _, elem := range existingChans {
		// Skip channels that are not categories
		if elem.Type != discordgo.ChannelTypeGuildCategory {
			continue
		}
		if strings.EqualFold(content, elem.Name) {
			bot.log.Warn("Category already exists!")
			if err = bot.SendMessage(catExistsMessage, bot.buildChannelID, session); err != nil {
				bot.log.WithError(err).Error("Error sending message")
			}
			return
		}
	}

	// Create the category first
	category, err := session.GuildChannelCreate(bot.guildID, content, discordgo.ChannelTypeGuildCategory)
	if err != nil {
		bot.log.WithError(err).WithField("userID", message.Author.ID).Error("Could not create category.")
		if err = bot.SendMessage(errorMessage, bot.buildChannelID, session); err != nil {
			bot.log.WithError(err).Error("Error sending message")
		}
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
			bot.log.WithError(err).WithField("userID", message.Author.ID).WithField("channelName", chanName).Error("Could not create the channel.")
			if err = bot.SendMessage(errorMessage, bot.buildChannelID, session); err != nil {
				bot.log.WithError(err).Error("Error sending message")
			}
			return
		}
	}

	// Send back a message confirming everything has gone as expected
	if err = bot.SendMessage(doneMessage, bot.buildChannelID, session); err != nil {
		bot.log.WithError(err).Error("Error sending message.")
		return
	}
}
