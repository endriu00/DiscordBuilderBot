package handler

import (
	"github.com/bwmarrin/discordgo"
)

// `UserJoinHandler` handles the situation in which a user joins the server.
func (bot *_bot) UserJoinHandler(session *discordgo.Session, event discordgo.GuildMemberAdd) {
	userID := event.Member.User.ID
	username := event.Member.User.Username

	// Insert the user in the database
	if err := bot.db.AddUser(userID, username); err != nil {
		bot.log.WithError(err).WithField("userID", userID).Error("Could not add user. Maybe it has already been inserted in the database?")
		return
	}
}
