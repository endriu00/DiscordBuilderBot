package service

import (
	"github.com/bwmarrin/discordgo"
)

func (bot *_bot) MessageReceivedCountHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	// update the points the user has in the database
	userID := message.Author.ID
	if err := bot.db.UpdateUserPoints(userID, 1); err != nil {
		bot.log.WithError(err).WithField("userID", userID).Error("Could not update user points.")
		return
	}
	bot.db.GetUserPoints()

	// if user has enough points he will be promoted
	// maybe change this in:
	// user asks for promotion
	// if he has enough points, he gets the promotion
	// if not, the points he has are returned

}
