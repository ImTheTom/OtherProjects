package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ImTheTom/BuggerLogger/config"
	"github.com/ImTheTom/BuggerLogger/db"
	"github.com/sirupsen/logrus"
)

const sleepTime = 5

var (
	errInvalidConfig = errors.New("Config failed to init correctly. Check previous logs")
	errInvalidLogger = errors.New("Logger failed to init correctly. Check previous logs")
)

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	if !config.IsLoggerReady() {
		return errInvalidLogger
	}

	if !config.IsConfigSetup() {
		return errInvalidConfig
	}

	logrus.Info("Logger and config setup correctly. Continuing to start the program")

	// Sleeping just before connecting to the DB to give it some time between starts and stops
	time.Sleep(sleepTime * time.Second)

	if err := db.CreateBuggerDB(config.GetConfig().DatabaseConnection); err != nil {
		return fmt.Errorf("Failed to connect to the db %w", err)
	}

	logrus.Info("Successfully connected to the database")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	logrus.Info("Exit signal reached. Shutting down the bugger logger.")

	db.CloseDatabase()

	return nil
}
