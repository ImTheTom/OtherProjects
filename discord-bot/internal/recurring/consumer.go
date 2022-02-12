package recurring

import (
	"context"

	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/sirupsen/logrus"
)

func processSyncUsers(syncUserCha <-chan model.User) {
	logrus.Info("Starting sync users channel consumption...")

	ctx := context.Background()

	for {
		user := <-syncUserCha

		if err := DBInt.UpsertUser(ctx, user); err != nil {
			logrus.Errorf("DB Upsert failed %v", err)
		}
	}
}

func processIncreasePoints(increasePointsCha <-chan model.User) {
	logrus.Info("Starting increase users channel consumption...")

	ctx := context.Background()

	for {
		user := <-increasePointsCha

		if err := DBInt.IncreasePoints(ctx, user); err != nil {
			logrus.Errorf("DB Upsert failed %v", err)
		}
	}
}
