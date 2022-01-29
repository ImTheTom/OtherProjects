package model

import "time"

type Gamble struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	GuildID   string    `json:"guild_id"`
	Amount    int       `json:"amount"`
	Winner    bool      `json:"winner"`
	CreatedAt time.Time `json:"created_at"`
}
