package db

import (
	"context"
	"errors"

	"github.com/ImTheTom/OtherProjects/discord-bot/model"
)

var errNoDb = errors.New("No database")

func (disDB discordDB) UpsertUser(ctx context.Context, user model.User) error {
	if disDB.db == nil {
		return errNoDb
	}

	disDB.mu.Lock()

	_, err := disDB.db.Exec(
		ctx,
		"INSERT INTO users(user_id,guild_id,username,nickname) VALUES ($1,$2,$3,$4) "+
			"ON CONFLICT (user_id, guild_id) DO UPDATE SET username=$3, nickname=$4",
		user.UserID,
		user.GuildID,
		user.Username,
		user.Nickname,
	)

	disDB.mu.Unlock()

	return err
}
