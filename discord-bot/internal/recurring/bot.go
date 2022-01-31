package recurring

import (
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/bot"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/db"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/helper"
	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

const (
	start = "0"
	limit = 1000
)

func SyncUsers() {
	logrus.Info("Syncing users now")

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
			logrus.Errorf("Errored getting guild members %v", err)

			continue
		}

		for _, mem := range members {
			user := model.User{
				UserID:   mem.User.ID,
				GuildID:  v.ID,
				Username: mem.User.Username,
				Nickname: mem.Nick,
			}

			err = db.GetDatabaseInterface().UpsertUser(helper.CreateContextWithTimeout(), user)
			if err != nil {
				logrus.Errorf("DB Upsert failed %v", err)
			}
		}
	}
}

func IncreasePoints() {
	logrus.Info("Increasing user points now")

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

				err := db.GetDatabaseInterface().IncreasePoints(helper.CreateContextWithTimeout(), user)
				if err != nil {
					logrus.Errorf("DB increase failed %v", err)
				}
			}
		}
	}
}
