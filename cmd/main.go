package main

import (
	"fmt"
	dgo "github.com/bwmarrin/discordgo"
	"github.com/endriu00/DiscordBuilderBot/service"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}

// `run` runs the application.
func run() error {
	// Start loggers. One is for bot, one for system communications
	log := logrus.NewEntry(logrus.StandardLogger())
	sysLog := logrus.NewEntry(logrus.StandardLogger())

	// Load configuration
	cfg, err := loadConfigurationFromEnv()
	if err != nil {
		sysLog.Error("Failed loading configuration.")
		return err
	}

	// Create bot
	bot := service.New(service.Config{
		Log:              log,
		GuildID:          cfg.GuildID,
		BuildChannelID:   cfg.BuildChannelID,
		ChannelsToCreate: cfg.ChannelsToCreate,
	})

	// Start a new discord session
	session, err := dgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		sysLog.WithError(err).Error("Could not create a new session.")
		return err
	}

	// Add handler messageHandler
	handlerRemover := session.AddHandler(bot.MessageBuildCategoryHandler)
	defer handlerRemover()

	// Open a websocket towards Discord
	if err = session.Open(); err != nil {
		sysLog.WithError(err).Error("Could not open a new websocket.")
		return err
	}
	defer func() {
		err = session.Close()
		if err != nil {
			sysLog.WithError(err).Error("Could not close the websocket.")
		}
	}()

	// Make channel for receiving signals.
	// The channel blocks the execution of the function.
	signalChan := make(chan os.Signal, 1)
	sysLog.Info("Bot is listening.")
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-signalChan

	return nil
}
