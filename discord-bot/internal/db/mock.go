package db

import (
	"context"

	"github.com/ImTheTom/OtherProjects/discord-bot/model"
)

type MockDiscordDBInterface struct {
	FakeUpsertUser                func(ctx context.Context, user model.User) error
	FakeInsertGamble              func(ctx context.Context, gamble model.Gamble) error
	FakeFindLatestGambleForUser   func(ctx context.Context, user model.User) (model.Gamble, error)
	FakeIncreasePoints            func(ctx context.Context, user model.User) error
	FakeSetUserPoints             func(ctx context.Context, user model.User) error
	FakeFindByUserIDAndGuildID    func(ctx context.Context, userID, guildID string) (model.User, error)
	FakeFindTopTenPointsForAGuild func(ctx context.Context, guildID string) ([]model.User, error)
}

func (m MockDiscordDBInterface) UpsertUser(ctx context.Context, user model.User) error {
	return nil
}

func (m MockDiscordDBInterface) InsertGamble(ctx context.Context, gamble model.Gamble) error {
	return nil
}

func (m MockDiscordDBInterface) FindLatestGambleForUser(ctx context.Context, user model.User) (model.Gamble, error) {
	return model.Gamble{}, nil
}

func (m MockDiscordDBInterface) IncreasePoints(ctx context.Context, user model.User) error {
	return nil
}

func (m MockDiscordDBInterface) SetUserPoints(ctx context.Context, user model.User) error {
	return nil
}

func (m MockDiscordDBInterface) FindByUserIDAndGuildID(
	ctx context.Context, userID, guildID string,
) (model.User, error) {
	return model.User{}, nil
}

func (m MockDiscordDBInterface) FindTopTenPointsForAGuild(ctx context.Context, guildID string) ([]model.User, error) {
	return []model.User{}, nil
}
