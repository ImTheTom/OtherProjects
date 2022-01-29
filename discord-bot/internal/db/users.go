package db

import (
	"context"
	"fmt"

	"github.com/ImTheTom/OtherProjects/discord-bot/model"
)

func FindByUserIDAndGuildID(ctx context.Context, userID, guildID string) (model.User, error) {
	user := model.User{}
	db := GetDatabase()

	fmt.Printf("USER - %s Guild %s\n", userID, guildID)

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