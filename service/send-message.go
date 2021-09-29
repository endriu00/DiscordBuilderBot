package service

import (
	"github.com/bwmarrin/discordgo"
)

func (bot *_bot) SendMessage(message, channelID string, session *discordgo.Session) error {
	_, err := session.ChannelMessageSend(channelID, message)
	if err != nil {
		bot.log.WithError(err).Error("Could not send a message.")
		return err
	}
	return nil
}
