package recurring

import (
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/bot"
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/db"
	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

const (
	start = "0"
	limit = 1000
)

var DBInt db.DiscordDBInterface

// Send only channel.
func syncUsers(syncUserCha chan<- model.User) bool {
	logrus.Info("Syncing users now")

	session := bot.GetSession()
	if session == nil {
		return false
	}

	state := session.State
	if state == nil {
		return false
	}

	guilds := state.Guilds

	for _, v := range guilds {
		members, err := session.GuildMembers(v.ID, start, limit)
		if err != nil {
			logrus.Errorf("Errored getting guild members %v", err)

			continue
		}

		for _, mem := range members {
			// I don't care about the bot itself
			if mem.User.ID == state.User.ID {
				continue
			}

			user := model.User{
				UserID:   mem.User.ID,
				GuildID:  v.ID,
				Username: mem.User.Username,
				Nickname: mem.Nick,
			}

			syncUserCha <- user
		}
	}

	return true
}

func increasePoints(increasePointsCha chan<- model.User) bool {
	logrus.Info("Increasing user points now")

	session := bot.GetSession()
	if session == nil {
		return false
	}

	state := session.State
	if state == nil {
		return false
	}

	guilds := state.Guilds

	for _, v := range guilds {
		guildPresences := v.Presences

		for _, pres := range guildPresences {
			if pres.Status != discordgo.StatusOffline {
				// I don't care about the bot itself
				if pres.User.ID == state.User.ID {
					continue
				}

				user := model.User{
					UserID:  pres.User.ID,
					GuildID: v.ID,
				}

				increasePointsCha <- user
			}
		}
	}

	return true
}
