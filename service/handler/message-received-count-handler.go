package handler

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
)

// `MessageReceivedCountHandler` handles the event of a user sending a message to the server.
// It updates user points of a pre-established amount, checks if it is necessary to promote
// the user, and, in that case, it sends a message to the server notifying him.
func (bot *_bot) MessageReceivedCountHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	//If the author of the message was the bot itself
	if message.Author.ID == session.State.User.ID {
		return
	}
	// Update the points the user has
	userID := message.Author.ID
	if err := bot.db.UpdateUserPoints(userID, 1); err != nil {
		bot.log.WithError(err).WithField("userID", userID).Error("Could not update user points.")
		return
	}

	// Get user points
	points, err := bot.db.GetUserPoints(userID)
	if err == sql.ErrNoRows {
		bot.log.WithField("userID", userID).Warn("User is not registered!")
		// TO BE CHANGED. USED THIS TO DO A SEAMLESS CREATION OF THE DB.
		err = bot.db.AddUser(userID, message.Author.Username)
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

	nextRole, err := bot.db.GetUserNextRole(userID)
	if err == sql.ErrNoRows {
		bot.log.Warn("Cannot promote again. Highest level.")
		return
	}
	if err != nil {
		bot.log.WithError(err).WithField("userID", userID).Error("Could not get next role.")
		return
	}
	// If the user has enough points, promote him
	if points >= nextRole.MinPoints {
		if err = bot.db.AddUserRole(userID, nextRole.ID); err != nil {
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
