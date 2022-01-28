package bot

import (
	"fmt"

	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

var sess *discordgo.Session

func Init() {
	fmt.Println("Initialising bot and commands")

	botToken := config.GetConfig().BotToken

	session, err := discordgo.New("Bot " + botToken)
	if err != nil {
		panic(err)
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	session.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = session.Open()
	if err != nil {
		panic(err)
	}

	sess = session
}

func GetSession() *discordgo.Session {
	if sess != nil {
		return sess
	}

	Init()

	return sess
}

func CloseBot() {
	if sess != nil {
		sess.Close()
	}
}
