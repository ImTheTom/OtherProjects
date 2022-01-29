package db

import (
	"context"

	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/sirupsen/logrus"
)

func FindByUserIDAndGuildID(ctx context.Context, userID, guildID string) (model.User, error) {
	user := model.User{}
	db := GetDatabase()

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

func FindTopTenPointsForAGuild(ctx context.Context, guildID string) ([]model.User, error) {
	users := make([]model.User, 0)
	db := GetDatabase()

	rows, err := db.Query(
		ctx,
		`SELECT * FROM users where guild_id = $1
		ORDER BY points DESC LIMIT 10`,
		guildID,
	)
	if err != nil {
		return users, err
	}

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

		users = append(users, usr)
	}

	return users, err
}
