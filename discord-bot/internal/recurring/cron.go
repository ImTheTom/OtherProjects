package recurring

import (
	"github.com/ImTheTom/OtherProjects/discord-bot/internal/db"
	"github.com/ImTheTom/OtherProjects/discord-bot/model"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

const chanSize = 1

var (
	cro                *cron.Cron
	increasePointsChan chan model.User
	syncUserChan       chan model.User
)

func Init() {
	increasePointsChan = make(chan model.User, chanSize)
	syncUserChan = make(chan model.User, chanSize)

	c := cron.New()
	if _, err := c.AddFunc("@every 3m", func() { syncUsers(syncUserChan) }); err != nil {
		logrus.Fatalf("Failed to add cron function, restarting... %v", err)
	}

	if _, err := c.AddFunc("@every 1m", func() { increasePoints(increasePointsChan) }); err != nil {
		logrus.Fatalf("Failed to add cron function, restarting... %v", err)
	}

	DBInt = db.GetDatabaseInterface()

	c.Start()

	logrus.Info("Crons have started")

	cro = c

	logrus.Info("Starting consumers")

	go processSyncUsers(syncUserChan)
	go processIncreasePoints(increasePointsChan)
}

func Stop() {
	cro.Stop()
	close(increasePointsChan)
	close(syncUserChan)
}
