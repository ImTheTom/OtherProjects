package main

import (
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

	if err := config.Init(); err != nil {
		logrus.Fatalf("Config failed to init, restarting... %v", err)
	}

	config.InitLogger()

	time.Sleep(sleepTime * time.Second)

	if _, err := db.NewDiscordDB(config.GetConfig().DatabaseConnection); err != nil {
		logrus.Fatalf("Failed to connect to the db, restarting... %v", err)
	}

	go recurring.Init()

	go bot.Init()

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	logrus.Info("Exit signal reached. Shutting down the bot.")
	bot.CloseBot()
	db.CloseDatabase()
	recurring.Stop()
}
