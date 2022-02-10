package db

import (
	"context"
	"errors"

	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/jackc/pgx"
)

func (disDB discordDB) InsertGamble(ctx context.Context, gamble model.Gamble) error {
	if disDB.db == nil {
		return errNoDb
	}

	disDB.mu.Lock()

	_, err := disDB.db.Exec(
		ctx,
		"INSERT INTO gambles(user_id,guild_id,amount,winner,created_at) VALUES ($1,$2,$3,$4,$5)",
		gamble.UserID,
		gamble.GuildID,
		gamble.Amount,
		gamble.Winner,
		gamble.CreatedAt,
	)

	disDB.mu.Unlock()

	return err
}

func (disDB discordDB) FindLatestGambleForUser(ctx context.Context, user model.User) (model.Gamble, error) {
	gamble := model.Gamble{}

	err := disDB.db.QueryRow(
		ctx,
		"SELECT id, user_id,guild_id,amount,winner,created_at FROM gambles "+
			"WHERE user_id = $1 AND guild_id = $2 ORDER BY id DESC LIMIT 1",
		user.UserID,
		user.GuildID,
	).Scan(
		&gamble.ID, &gamble.UserID, &gamble.GuildID, &gamble.Amount, &gamble.Winner, &gamble.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return gamble, nil
	}

	return gamble, err
}
