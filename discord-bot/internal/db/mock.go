package db

import (
	"context"
	"errors"
	"time"

	"github.com/ImTheTom/OtherProjects/discord-bot/model"
)

var errStand = errors.New("Standard DB fail")

type MockDiscordDBInterface struct{}

func (m MockDiscordDBInterface) UpsertUser(ctx context.Context, user model.User) error {
	if user.UserID == "0" {
		return errStand
	}

	return nil
}

func (m MockDiscordDBInterface) InsertGamble(ctx context.Context, gamble model.Gamble) error {
	if gamble.Amount == 1 {
		return errStand
	}

	return nil
}

func (m MockDiscordDBInterface) FindLatestGambleForUser(ctx context.Context, user model.User) (model.Gamble, error) {
	if user.UserID == "0" {
		return model.Gamble{}, errStand
	}

	if user.UserID == "1" {
		return model.Gamble{
			ID:        1,
			UserID:    "1",
			GuildID:   "1",
			Amount:    1,
			Winner:    true,
			CreatedAt: time.Now(),
		}, nil
	}

	return model.Gamble{
		ID:        1,
		UserID:    "1",
		GuildID:   "1",
		Amount:    1,
		Winner:    true,
		CreatedAt: time.Now().AddDate(0, 0, -1),
	}, nil
}

func (m MockDiscordDBInterface) IncreasePoints(ctx context.Context, user model.User) error {
	if user.UserID == "0" {
		return errStand
	}

	return nil
}

func (m MockDiscordDBInterface) SetUserPoints(ctx context.Context, user model.User) error {
	if user.UserID == "0" {
		return errStand
	}

	return nil
}

func (m MockDiscordDBInterface) FindByUserIDAndGuildID(
	ctx context.Context, userID, guildID string,
) (model.User, error) {
	if userID == "2" {
		return model.User{}, errStand
	}

	return model.User{}, nil
}

func (m MockDiscordDBInterface) FindTopTenPointsForAGuild(ctx context.Context, guildID string) ([]model.User, error) {
	if guildID == "1" {
		return []model.User{}, errStand
	}

	return []model.User{}, nil
}
