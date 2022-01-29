package recurring

import (
	"github.com/robfig/cron/v3"
)

var cro *cron.Cron

func Init() {
	c := cron.New()
	if _, err := c.AddFunc("@every 3m", func() { syncUsers() }); err != nil {
		panic(err)
	}

	if _, err := c.AddFunc("@every 1m", func() { increasePoints() }); err != nil {
		panic(err)
	}

	c.Start()
	cro = c

	for {
		select {}
	}
}

func Stop() {
	cro.Stop()
}
