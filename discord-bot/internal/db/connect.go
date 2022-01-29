package db

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	retries   = 50
	sleepTime = 5
)

var (
	dbURL              string
	db                 *pgxpool.Pool
	errFailedToConnect = errors.New("Failed to connect to db")
)

func Connect(url string) error {
	dbURL = url

	var pool *pgxpool.Pool

	var err error

	for i := 0; i < retries; i++ {
		pool, err = pgxpool.Connect(context.Background(), dbURL)
		if err != nil {
			logrus.Errorf("Db connect errored %v", err)
			time.Sleep(sleepTime * time.Second)

			continue
		}

		break
	}

	if err != nil {
		return errFailedToConnect
	}

	logrus.Info("Successfully connected to the database")

	db = pool

	return nil
}

func GetDatabase() *pgxpool.Pool {
	if db == nil {
		if err := Connect(dbURL); err != nil {
			logrus.Errorf("Db connect errored %v", err)
		}
	}

	return db
}

func CloseDatabase() {
	if db != nil {
		db.Close()
	}
}
