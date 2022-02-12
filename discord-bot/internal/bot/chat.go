package bot

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

const (
	birthCommand  = "birth"
	flipCommand   = "flip"
	helloCommand  = "hello"
	joinedCommand = "joined"
	louisCommand  = "louis"
	pingCommand   = "ping"
	rapCommand    = "rap"

	randomUpperLimit = 2
	hoursInADay      = 24
	layoutUS         = "January 2, 2006"
)

var louisQuotes = []string{
	"Trying to solo carry",
	"I carried",
	"Give me my blue buff",
	"Dont ks my penta",
	"Honour me",
	"I solo carried btw",
	"I would carry this game so hard if only I wasn't lagging.",
	"That champion is so lame",
	"Stop taking my cs",
	"I didn't crit?",
	"I don't care about honours btw.",
}

func standardChatMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	prefix := config.GetConfig().Prefix

	switch m.Content {
	case fmt.Sprintf("%s%s", prefix, birthCommand):
		birthMessage(s, m)
	case fmt.Sprintf("%s%s", prefix, flipCommand):
		flipMessage(s, m)
	case fmt.Sprintf("%s%s", prefix, helloCommand):
		helloMessage(s, m)
	case fmt.Sprintf("%s%s", prefix, joinedCommand):
		joinedMessage(s, m)
	case fmt.Sprintf("%s%s", prefix, louisCommand):
		louisMessage(s, m)
	case fmt.Sprintf("%s%s", prefix, pingCommand):
		pongMessage(s, m)
	case fmt.Sprintf("%s%s", prefix, rapCommand):
		rapMessage(s, m)
	}
}

func birthMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	logMessage(m)

	created := time.Date(2017, 5, 18, 12, 0, 0, 0, time.UTC)
	now := time.Now()

	days := int(now.Sub(created).Hours() / hoursInADay)

	message := fmt.Sprintf("I was created on the 18th of May 2017. This means I'm %d days old!", days)

	communicateStandardMessage(s, m, message)
}

func joinedMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	logMessage(m)

	mem, err := s.GuildMember(m.GuildID, m.Author.ID)
	if err != nil {
		return
	}

	joined, err := mem.JoinedAt.Parse()
	if err != nil {
		return
	}

	now := time.Now()

	days := int(now.Sub(joined).Hours() / hoursInADay)

	message := fmt.Sprintf("You joined on %v. That was %d days ago!", joined.Format(layoutUS), days)

	communicateStandardMessage(s, m, message)
}

func flipMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	logMessage(m)

	message := "Heads"

	if value := rand.Intn(randomUpperLimit); value == 0 {
		message = "Tails"
	}

	communicateStandardMessage(s, m, message)
}

func helloMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	logMessage(m)

	helloString := fmt.Sprintf("Hello %s", m.Author.Mention())
	communicateStandardMessage(s, m, helloString)
}

func louisMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	logMessage(m)

	index := rand.Intn(len(louisQuotes) - 1)
	communicateStandardMessage(s, m, louisQuotes[index])
}

func pongMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	logMessage(m)

	communicateStandardMessage(s, m, "Pong!")
}

func rapMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	logMessage(m)

	rapString := "Got rocks on my wrist, that shit you can't resist."
	communicateStandardMessage(s, m, rapString)
}
