package main

import (
	"context"
	"fmt"
	dgo "github.com/bwmarrin/discordgo"
	"github.com/endriu00/DiscordBuilderBot/service/db"
	"github.com/endriu00/DiscordBuilderBot/service/handler"
	pgx "github.com/jackc/pgx/v4/pgxpool"
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

	// Define an empty context. Can be managed and extended in future versions
	ctx := context.Background()

	// Connect to the database and ping it
	database, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		sysLog.WithError(err).Error("Failed to connect to DB")
		return err
	}
	if err = database.Ping(ctx); err != nil {
		sysLog.WithError(err).Error("Database is not responding")
		return err
	}

	// Initialize the database for the bot
	botDB, err := db.New(database)
	if err != nil {
		sysLog.WithError(err).Error("Failed to create the database for the bot.")
		return err
	}

	// Create bot
	bot := handler.New(handler.Config{
		Log:                log,
		GuildID:            cfg.GuildID,
		BuildChannelID:     cfg.BuildChannelID,
		GetPointsChannelID: cfg.GetPointsChannelID,
		ChannelsToCreate:   cfg.ChannelsToCreate,
		DB:                 botDB,
	})

	// Start a new discord session
	session, err := dgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		sysLog.WithError(err).Error("Could not create a new session.")
		return err
	}

	// Add message creation handlers
	handlerRemoverBuild := session.AddHandler(bot.MessageBuildCategoryHandler)
	defer handlerRemoverBuild()
	handlerRemoverCount := session.AddHandler(bot.MessageReceivedCountHandler)
	defer handlerRemoverCount()
	handlerRemoverPoints := session.AddHandler(bot.MessageGetPointsHandler)
	defer handlerRemoverPoints()

	// Add user joining the server handlers
	handlerRemoverJoin := session.AddHandler(bot.UserJoinHandler)
	defer handlerRemoverJoin()

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
	// It stays blocked until it receives some system signals.
	signalChan := make(chan os.Signal, 1)
	sysLog.Info("Bot is listening.")
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-signalChan

	return nil
}
