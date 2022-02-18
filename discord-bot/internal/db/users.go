package db

import (
	"context"

	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/sirupsen/logrus"
)

const userListSize = 10

func (disDB *discordDB) FindByUserIDAndGuildID(ctx context.Context, userID, guildID string) (model.User, error) {
	user := model.User{}
	db := disDB.db

	err := db.QueryRow(
		ctx,
		"SELECT user_id,guild_id,username,nickname,points FROM users WHERE user_id = $1 AND guild_id = $2",
		userID,
		guildID,
	).Scan(
		&user.UserID, &user.GuildID, &user.Username, &user.Nickname, &user.Points,
	)

	return user, err
}

func (disDB *discordDB) FindTopTenPointsForAGuild(ctx context.Context, guildID string) ([]model.User, error) {
	users := make([]model.User, userListSize)
	db := disDB.db

	rows, err := db.Query(
		ctx,
		`SELECT * FROM users where guild_id = $1
		ORDER BY points DESC LIMIT 10`,
		guildID,
	)
	if err != nil {
		return users, err
	}

	i := 0

	for rows.Next() {
		var usr model.User

		err := rows.Scan(
			&usr.UserID,
			&usr.GuildID,
			&usr.Username,
			&usr.Nickname,
			&usr.Points,
		)
		if err != nil {
			logrus.Errorf("Scan failed %v", err)

			continue
		}

		users[i] = usr
		i++
	}

	return users, err
}
