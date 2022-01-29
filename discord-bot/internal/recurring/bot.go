package recurring

import (
	"fmt"

	"github.com/ImTheTom/OtherProjects/discord-bot/internal/bot"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/db"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/helper"
	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/bwmarrin/discordgo"
)

const (
	start = "0"
	limit = 1000
)

func syncUsers() {
	fmt.Println("Syncing users now")

	session := bot.GetSession()
	if session == nil {
		return
	}

	state := session.State
	if state == nil {
		return
	}

	guilds := state.Guilds

	for _, v := range guilds {
		members, err := session.GuildMembers(v.ID, start, limit)
		if err != nil {
			fmt.Printf("error is %v\n", err)

			continue
		}

		for _, mem := range members {
			user := model.User{
				UserID:   mem.User.ID,
				GuildID:  v.ID,
				Username: mem.User.Username,
				Nickname: mem.Nick,
			}

			err = db.UpsertUser(helper.CreateContextWithTimeout(), user)
			if err != nil {
				fmt.Printf("Error is %v\n", err)
			}
		}
	}
}

func increasePoints() {
	fmt.Println("Increasing user points now")

	session := bot.GetSession()
	if session == nil {
		return
	}

	state := session.State
	if state == nil {
		return
	}

	guilds := state.Guilds

	for _, v := range guilds {
		guildPresences := v.Presences

		for _, pres := range guildPresences {
			if pres.Status != discordgo.StatusOffline {
				user := model.User{
					UserID:  pres.User.ID,
					GuildID: v.ID,
				}

				err := db.IncreasePoints(helper.CreateContextWithTimeout(), user)
				if err != nil {
					fmt.Printf("Error is %v\n", err)
				}
			}
		}
	}
}
