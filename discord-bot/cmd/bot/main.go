package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/bot"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/db"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/recurring"
	"github.com/sirupsen/logrus"
)

const sleepTime = 5

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	config.InitLogger()

	if err := config.Init(); err != nil {
		return fmt.Errorf("Config failed to init, restarting... %w", err)
	}

	logrus.Info("Finished configuration. Sleeping for 5 seconds before connecting to the db and discord...")

	time.Sleep(sleepTime * time.Second)

	if err := db.NewDiscordDB(config.GetConfig().DatabaseConnection); err != nil {
		return fmt.Errorf("Failed to connect to the db, restarting... %w", err)
	}

	logrus.Info("Running go routines")

	go recurring.Init()

	go bot.Init()

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	logrus.Info("Exit signal reached. Shutting down the bot.")
	bot.CloseBot()
	db.CloseDatabase()
	recurring.Stop()

	return nil
}
