package handler

import (
	"context"
	"github.com/bwmarrin/discordgo"
	pgx "github.com/jackc/pgx/v4"
)

// TODO: instead of checking for each message received the points the user has left
// to the next role, store in a in-memory buffer the latest amount of points received
// and update it instead of the database. Connect to the database every now and then
// in order to avoid stressing it. THINK ABOUT A SIMILAR SOLUTION.
// MAYBE REDIS?

// `MessageReceivedCountHandler` handles the event of a user sending a message to the server.
// It updates user points of a pre-established amount, checks if it is necessary to promote
// the user, and, in that case, it sends a message to the server notifying him.
func (bot *_bot) MessageReceivedCountHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	ctx := context.Background()
	userID := message.Author.ID
	pointsWorth := 0

	// If the author of the message was the bot itself
	if userID == session.State.User.ID {
		return
	}

	// TODO: Check if the message is a good message or a bad message

	// Update the points the user has considering the type of the message sent.
	// If the message does not belong to the types shown, do not consider it and return.
	switch message.Type {
	case discordgo.MessageTypeDefault:
		pointsWorth = messageSentPoints
	case discordgo.MessageTypeReply:
		pointsWorth = messageReplyPoints
	default:
		return
	}
	if err := bot.db.UpdateUserPoints(userID, pointsWorth, ctx); err != nil {
		bot.log.WithError(err).WithField("userID", userID).Error("Could not update user points.")
		return
	}

	// Get user points
	points, err := bot.db.GetUserPoints(userID, ctx)
	if err == pgx.ErrNoRows {
		bot.log.WithField("userID", userID).Warn("User is not registered!")
		// TO BE CHANGED. USED THIS TO DO A SEAMLESS CREATION OF THE DB.
		err = bot.db.AddUser(userID, message.Author.Username, ctx)
		if err != nil {
			bot.log.WithError(err).Error("Could not create the user.")
			return
		}
		return
	}
	if err != nil {
		bot.log.WithError(err).WithField("userID", userID).Error("Could not get user points.")
		return
	}

	// Get user next role for the server
	nextRole, err := bot.db.GetUserNextRole(userID, ctx)
	if err == pgx.ErrNoRows {
		bot.log.Warn("Cannot promote again. Highest level.")
		return
	}
	if err != nil {
		bot.log.WithError(err).WithField("userID", userID).Error("Could not get next role.")
		return
	}

	// If the user has enough points, promote him
	if points >= nextRole.MinPoints {
		if err = bot.db.AddUserRole(userID, nextRole.ID, ctx); err != nil {
			bot.log.WithError(err).WithField("userID", userID).Error("Could not upgrade user role on the database.")
			return
		}

		if err = session.GuildMemberRoleAdd(bot.guildID, userID, nextRole.ID); err != nil {
			bot.log.WithError(err).WithField("userID", userID).Error("Could not upgrade user role on the server.")
			return
		}
		if err = bot.SendMessage("Congratulation, "+message.Author.Username+"! You have been promoted to "+nextRole.Name+"!", message.ChannelID, session); err != nil {
			bot.log.WithError(err).WithField("userID", userID).Error("Could not send message.")
			return
		}
	}
}
