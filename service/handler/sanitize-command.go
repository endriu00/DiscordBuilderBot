package service

import (
	"strings"
	"text/scanner"
)

// `SanitizeCommand` takes the content of the received message and scans it.
// If it starts with the correct beginning character `!`, it is a good command.
// Otherwise, it is a command not meant for the bot.
// It takes as input the content of the message and returns the proper command
// and an error, if any.
func (bot *_bot) SanitizeCommand(content string) (string, error) {
	var stringScan scanner.Scanner

	// Initialize the string scanner for the content string
	stringScan.Init(strings.NewReader(content))

	// Scan the next token. It should be the character `!`.
	// Then convert the scanned token to a string.
	_ = stringScan.Scan()
	commandIdentifier := stringScan.TokenText()

	// If the token was not the correct character, return
	if commandIdentifier != "!" {
		bot.log.Warn("Message not meant to be a command for the bot.")
		return "", ErrNotCommand
	}

	// Scan the next token. It should be the command itself.
	_ = stringScan.Scan()
	command := stringScan.TokenText()
	return command, nil
}
