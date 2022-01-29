package bot

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/db"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/helper"
	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/bwmarrin/discordgo"
)

const (
	points      = "points"
	leaderBoard = "ladder"
	gamble      = "gamble"

	allPointsGambleWin   = 2
	numberOfGambleParams = 2
	allPointsGamble      = "all"
)

var errInvalidGambleAmount = errors.New("Invalid Gamble Amount")

func GambleInteractions(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	prefix := config.GetConfig().Prefix

	called := false

	switch m.Content {
	case fmt.Sprintf("%s%s", prefix, points):
		called = true

		pointsMessage(s, m)
	case fmt.Sprintf("%s%s", prefix, leaderBoard):
		called = true

		leaderBoardMessage(s, m)
	}

	if called {
		return
	}

	contentSplit := strings.Split(m.Content, " ")
	if len(contentSplit) != numberOfGambleParams {
		return
	}

	if contentSplit[0] == fmt.Sprintf("%s%s", prefix, gamble) {
		gamblePoints(s, m, contentSplit[1])
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
	users, err := db.FindTopTenPointsForAGuild(helper.CreateContextWithTimeout(), m.GuildID)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	totalMessage := "Current ladder is.\n"
	for i, v := range users {
		totalMessage += fmt.Sprintf("%d - %s - %d\n", i+1, v.Username, v.Points)
	}

	if _, err := s.ChannelMessageSend(m.ChannelID, totalMessage); err != nil {
		fmt.Printf("failed to send message %v\n", err)
	}
}

func gamblePoints(s *discordgo.Session, m *discordgo.MessageCreate, amountParam string) {
	user, err := db.FindByUserIDAndGuildID(helper.CreateContextWithTimeout(), m.Author.ID, m.GuildID)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if user.Points == 0 {
		if _, err := s.ChannelMessageSend(m.ChannelID, "You have no points"); err != nil {
			fmt.Printf("failed to send message %v\n", err)
		}

		return
	}

	winner := false

	var currentPoints int

	if value := rand.Intn(randomUpperLimit); value == 0 {
		winner = true
	}

	if amountParam == allPointsGamble {
		currentPoints = calulatePointsAll(user, winner)
	} else {
		currentPoints, err = calulatePointsLessThanAll(user, amountParam, winner)
		if err != nil {
			if _, err := s.ChannelMessageSend(m.ChannelID, "Invalid gamble amount"); err != nil {
				fmt.Printf("failed to send message %v\n", err)
			}

			return
		}
	}

	user.Points = currentPoints

	if err = db.SetUserPoints(helper.CreateContextWithTimeout(), user); err != nil {
		return
	}

	if winner {
		winnerMessage := fmt.Sprintf("You won %s! Total points is now %d.", user.Username, user.Points)
		if _, err := s.ChannelMessageSend(m.ChannelID, winnerMessage); err != nil {
			fmt.Printf("failed to send message %v\n", err)
		}
	} else {
		winnerMessage := fmt.Sprintf("You lost %s. Total points is now %d.", user.Username, user.Points)
		if _, err := s.ChannelMessageSend(m.ChannelID, winnerMessage); err != nil {
			fmt.Printf("failed to send message %v\n", err)
		}
	}
}

func calulatePointsAll(user model.User, winner bool) int {
	if winner {
		return user.Points * allPointsGambleWin
	}

	return 0
}

func calulatePointsLessThanAll(user model.User, amountParam string, winner bool) (int, error) {
	currentPoints := 0

	gambleAmount, err := strconv.Atoi(amountParam)
	if err != nil {
		return currentPoints, err
	}

	if gambleAmount > user.Points {
		return currentPoints, errInvalidGambleAmount
	}

	if winner {
		currentPoints = user.Points + gambleAmount
	} else {
		currentPoints = user.Points - gambleAmount
	}

	return currentPoints, nil
}
