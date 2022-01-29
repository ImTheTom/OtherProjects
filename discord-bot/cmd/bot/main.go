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
)

const sleepTime = 5

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if err := config.Init(); err != nil {
		panic(err)
	}

	config.InitLogger()

	time.Sleep(sleepTime * time.Second)

	if err := db.Connect(config.GetConfig().DatabaseConnection); err != nil {
		panic(err)
	}

	go recurring.Init()

	bot.Init()

	fmt.Println("Bot is now running.")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	bot.CloseBot()
	db.CloseDatabase()
	recurring.Stop()
}
