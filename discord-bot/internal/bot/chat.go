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

func StandardChatMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	created := time.Date(2017, 5, 18, 12, 0, 0, 0, time.UTC)
	now := time.Now()

	days := int(now.Sub(created).Hours() / hoursInADay)

	message := fmt.Sprintf("I was created on the 18th of May 2017. This means I'm %d days old!", days)

	if _, err := s.ChannelMessageSend(m.ChannelID, message); err != nil {
		fmt.Printf("failed to send message %v\n", err)
	}
}

func joinedMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
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

	if _, err := s.ChannelMessageSend(m.ChannelID, message); err != nil {
		fmt.Printf("failed to send message %v\n", err)
	}
}

func flipMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "Heads"

	if value := rand.Intn(randomUpperLimit); value == 0 {
		message = "Tails"
	}

	if _, err := s.ChannelMessageSend(m.ChannelID, message); err != nil {
		fmt.Printf("failed to send message %v\n", err)
	}
}

func helloMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	helloString := fmt.Sprintf("Hello %s", m.Author.Mention())
	if _, err := s.ChannelMessageSend(m.ChannelID, helloString); err != nil {
		fmt.Printf("failed to send message %v\n", err)
	}
}

func louisMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	index := rand.Intn(len(louisQuotes) - 1)
	if _, err := s.ChannelMessageSend(m.ChannelID, louisQuotes[index]); err != nil {
		fmt.Printf("failed to send message %v\n", err)
	}
}

func pongMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSend(m.ChannelID, "Pong!"); err != nil {
		fmt.Printf("failed to send message %v\n", err)
	}
}

func rapMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	rapString := "Got rocks on my wrist, that shit you can't resist."
	if _, err := s.ChannelMessageSend(m.ChannelID, rapString); err != nil {
		fmt.Printf("failed to send message %v\n", err)
	}
}
