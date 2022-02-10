package bot

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/db"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/helper"
	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

const (
	points      = "points"
	leaderBoard = "ladder"
	gamble      = "gamble"

	allPointsGambleWin   = 2
	numberOfGambleParams = 2
	allPointsGamble      = "all"
)

var (
	errInvalidGambleAmount = errors.New("Invalid Gamble Amount")
	errRateLimitGamble     = errors.New("Can't gamble so quickly")
	errNoPoints            = errors.New("You have no points")
)

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

		PointsMessage(s, m)
	case fmt.Sprintf("%s%s", prefix, leaderBoard):
		called = true

		LeaderBoardMessage(s, m)
	}

	if called {
		return
	}

	contentSplit := strings.Split(m.Content, " ")
	if len(contentSplit) != numberOfGambleParams {
		return
	}

	if contentSplit[0] == fmt.Sprintf("%s%s", prefix, gamble) {
		GamblePoints(s, m, contentSplit[1])
	}
}

func PointsMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	logMessage(m)

	ctx, cancel := helper.CreateContextWithTimeout()

	defer cancel()

	user, err := db.GetDatabaseInterface().FindByUserIDAndGuildID(
		ctx, m.Author.ID, m.GuildID,
	)
	if err != nil {
		logrus.Warnf("Failed to find user of %s %s db error - %v", m.Author.ID, m.GuildID, err)
	}

	mess := fmt.Sprintf("You have a total of %d points %s", user.Points, m.Author.Mention())

	communicateStandardMessage(s, m, mess)
}

func LeaderBoardMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	logMessage(m)

	ctx, cancel := helper.CreateContextWithTimeout()

	defer cancel()

	users, err := db.GetDatabaseInterface().FindTopTenPointsForAGuild(ctx, m.GuildID)
	if err != nil {
		logrus.Warnf("DB error - %v", err)
	}

	totalMessage := "Current ladder is.\n"
	for i, v := range users {
		totalMessage += fmt.Sprintf("%d - %s - %d\n", i+1, v.Username, v.Points)
	}

	communicateStandardMessage(s, m, totalMessage)
}

func GamblePoints(s *discordgo.Session, m *discordgo.MessageCreate, amountParam string) {
	logMessage(m)

	ctx, cancel := helper.CreateContextWithTimeout()

	defer cancel()

	user, err := CheckGambleIsSane(ctx, m)
	if err != nil {
		communicateStandardMessage(s, m, err.Error())

		return
	}

	winner := false

	var currentPoints int

	if value := rand.Intn(randomUpperLimit); value == 0 {
		winner = true
	}

	if amountParam == allPointsGamble {
		currentPoints = CalulatePointsAll(ctx, user, winner)
	} else {
		currentPoints, err = CalulatePointsLessThanAll(ctx, user, amountParam, winner)
		if err != nil {
			communicateStandardMessage(s, m, "Invalid gamble amount")

			return
		}
	}

	user.Points = currentPoints

	if err = DBInt.SetUserPoints(ctx, user); err != nil {
		return
	}

	if winner {
		winnerMessage := fmt.Sprintf("You won %s! Total points is now %d.", user.Username, user.Points)
		communicateStandardMessage(s, m, winnerMessage)
	} else {
		loserMessage := fmt.Sprintf("You lost %s. Total points is now %d.", user.Username, user.Points)
		communicateStandardMessage(s, m, loserMessage)
	}
}

func CheckGambleIsSane(ctx context.Context, m *discordgo.MessageCreate) (model.User, error) {
	user, err := DBInt.FindByUserIDAndGuildID(
		ctx, m.Author.ID, m.GuildID,
	)
	if err != nil {
		return user, err
	}

	gamble, err := DBInt.FindLatestGambleForUser(ctx, user)
	if err != nil {
		if err.Error() != "no rows in result set" {
			return user, err
		}
	}

	if gamble.ID != 0 {
		currentTime := time.Now()

		t3 := currentTime.Sub(gamble.CreatedAt)
		if t3 < time.Minute {
			return user, errRateLimitGamble
		}
	}

	if user.Points <= 0 {
		return user, errNoPoints
	}

	return user, nil
}

func CalulatePointsAll(ctx context.Context, user model.User, winner bool) int {
	_ = SaveGamble(ctx, user, user.Points, winner)

	if winner {
		return user.Points * allPointsGambleWin
	}

	return 0
}

func CalulatePointsLessThanAll(ctx context.Context, user model.User, amountParam string, winner bool) (int, error) {
	currentPoints := 0

	gambleAmount, err := strconv.Atoi(amountParam)
	if err != nil {
		return currentPoints, err
	}

	if gambleAmount > user.Points {
		return currentPoints, errInvalidGambleAmount
	}

	if gambleAmount <= 0 {
		return currentPoints, errInvalidGambleAmount
	}

	if winner {
		currentPoints = user.Points + gambleAmount
	} else {
		currentPoints = user.Points - gambleAmount
	}

	_ = SaveGamble(ctx, user, gambleAmount, winner)

	return currentPoints, nil
}

func SaveGamble(ctx context.Context, user model.User, amount int, winner bool) error {
	gm := model.Gamble{
		UserID:    user.UserID,
		GuildID:   user.GuildID,
		Amount:    amount,
		Winner:    winner,
		CreatedAt: time.Now(),
	}

	err := DBInt.InsertGamble(ctx, gm)
	if err != nil {
		logrus.Errorf("Insert gamble failed %v", err)
	}

	return err
}
