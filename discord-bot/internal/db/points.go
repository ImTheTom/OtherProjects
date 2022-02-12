package db

import (
	"context"

	"github.com/ImTheTom/OtherProjects/discord-bot/model"
)

func (disDB *discordDB) IncreasePoints(ctx context.Context, user model.User) error {
	currentDB := disDB.db
	if currentDB == nil {
		return errNoDb
	}

	disDB.mu.Lock()

	_, err := currentDB.Exec(
		ctx,
		"UPDATE users SET points = points + 1 WHERE user_id = $1 AND guild_id = $2",
		user.UserID,
		user.GuildID,
	)

	disDB.mu.Unlock()

	return err
}

func (disDB *discordDB) SetUserPoints(ctx context.Context, user model.User) error {
	currentDB := disDB.db
	if currentDB == nil {
		return errNoDb
	}

	disDB.mu.Lock()

	_, err := currentDB.Exec(
		ctx,
		"UPDATE users SET points = $1 WHERE user_id = $2 AND guild_id = $3",
		user.Points,
		user.UserID,
		user.GuildID,
	)

	disDB.mu.Unlock()

	return err
}
