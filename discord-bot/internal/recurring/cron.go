package recurring

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var cro *cron.Cron

func Init() {
	c := cron.New()
	if _, err := c.AddFunc("@every 3m", func() { SyncUsers() }); err != nil {
		logrus.Fatalf("Failed to add cron function, restarting... %v", err)
	}

	if _, err := c.AddFunc("@every 1m", func() { IncreasePoints() }); err != nil {
		logrus.Fatalf("Failed to add cron function, restarting... %v", err)
	}

	c.Start()

	logrus.Info("Crons have started")

	cro = c
}

func Stop() {
	cro.Stop()
}
