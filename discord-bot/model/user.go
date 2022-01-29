package model

type User struct {
	UserID   string `json:"user_id"`
	GuildID  string `json:"guild_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Points   int    `json:"points"`
}
