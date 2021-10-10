package handler

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

// `MessageGetPointsHandler` is the handler for user asking for the amount of
// points he has reached.
func (bot *_bot) MessageGetPointsHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	userID := message.Author.ID
	ctx := context.Background()

	// Check whether the message came to the correct channel
	// or the userID was the bot ID return
	if message.ChannelID != bot.getPointsChannelID || message.Author.ID == session.State.User.ID {
		return
	}

	// Sanitize Input
	content, err := bot.SanitizeCommand(message.Content)
	if err == ErrNotCommand {
		bot.log.WithError(err).WithField("userID", message.Author.ID).Error("Not a command meant for the bot.")
		if err = bot.SendMessage(notCommandMessage, bot.getPointsChannelID, session); err != nil {
			bot.log.WithError(err).Error("Error sending message")
		}
		return
	}

	// Check whether message content is equal to "points"
	if !strings.EqualFold(content, "points") {
		bot.log.Warn("Command not recognized")
		if err = bot.SendMessage(notGetPointsMessage, bot.getPointsChannelID, session); err != nil {
			bot.log.WithError(err).Error("Error sending message")
		}
		return
	}

	// Get user points
	points, err := bot.db.GetUserPoints(userID, ctx)
	if err != nil {
		bot.log.WithError(err).WithField("userID", userID).Error("Could not get user points")
		if err = bot.SendMessage(errorMessage, bot.getPointsChannelID, session); err != nil {
			bot.log.WithError(err).Error("Error sending message")
		}
		return
	}

	// Send message
	if err = bot.SendMessage(pointsMessage+strconv.Itoa(points), bot.getPointsChannelID, session); err != nil {
		bot.log.WithError(err).WithField("userID", userID).Error("Could not send a message")
		return
	}
}
