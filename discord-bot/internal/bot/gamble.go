package bot

import (
	"fmt"
	"time"

	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/db"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/helper"
	"github.com/bwmarrin/discordgo"
)

const (
	points      = "points"
	leaderBoard = "ladder"
)

func GambleInteractions(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	prefix := config.GetConfig().Prefix

	switch m.Content {
	case fmt.Sprintf("%s%s", prefix, points):
		pointsMessage(s, m)
	case fmt.Sprintf("%s%s", prefix, leaderBoard):
		leaderBoardMessage(s, m)
	}
}

func pointsMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	user, err := db.FindByUserIDAndGuildID(helper.CreateContextWithTimeout(), m.Author.ID, m.GuildID)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	mess := fmt.Sprintf("You have a total of %d points %s", user.Points, m.Author.Mention())

	if _, err := s.ChannelMessageSend(m.ChannelID, mess); err != nil {
		fmt.Printf("failed to send message %v\n", err)
	}
}

func leaderBoardMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
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
