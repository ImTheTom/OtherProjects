package recurring

import (
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/helper"
	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/sirupsen/logrus"
)

func ProcessSyncUsers(syncUserCha <-chan model.User) {
	logrus.Info("Starting sync users channel consumption...")

	for {
		user := <-syncUserCha

		logrus.Infof("syncing user %+v", user)

		if err := DBInt.UpsertUser(helper.CreateContextWithTimeout(), user); err != nil {
			logrus.Errorf("DB Upsert failed %v", err)
		}
	}
}

func ProcessIncreasePoints(increasePointsCha <-chan model.User) {
	logrus.Info("Starting increase users channel consumption...")

	for {
		user := <-increasePointsCha

		logrus.Infof("Increasing points for user %+v", user)

		if err := DBInt.IncreasePoints(helper.CreateContextWithTimeout(), user); err != nil {
			logrus.Errorf("DB Upsert failed %v", err)
		}
	}
}
