package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/bot"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/db"
)

const sleepTime = 5

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	time.Sleep(sleepTime * time.Second)

	if err := db.Connect(config.GetConfig().DatabaseConnection); err != nil {
		panic(err)
	}

	bot.Init()

	fmt.Println("Bot is now running.")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	bot.CloseBot()
}
