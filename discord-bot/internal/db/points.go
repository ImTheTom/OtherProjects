package db

import (
	"context"

	"github.com/ImTheTom/OtherProjects/discord-bot/model"
)

func IncreasePoints(ctx context.Context, user model.User) error {
	currentDB := GetDatabase()
	if currentDB == nil {
		return errNoDb
	}

	_, err := currentDB.Exec(
		ctx,
		"UPDATE users SET points = points + 1 WHERE user_id = $1 AND guild_id = $2",
		user.UserID,
		user.GuildID,
	)

	return err
}

func SetUserPoints(ctx context.Context, user model.User) error {
	currentDB := GetDatabase()
	if currentDB == nil {
		return errNoDb
	}

	_, err := currentDB.Exec(
		ctx,
		"UPDATE users SET points = $1 WHERE user_id = $2 AND guild_id = $3",
		user.Points,
		user.UserID,
		user.GuildID,
	)

	return err
}
