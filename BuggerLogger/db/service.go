package db

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	retries   = 50
	sleepTime = 5
)

var errFailedToConnect = errors.New("Failed to connect to db")

var (
	dbConnOnce sync.Once
	_buggerDB  *buggerDB
)

type buggerDB struct {
	db *pgxpool.Pool
}

func CreateBuggerDB(connection string) error {
	var err error

	dbConnOnce.Do(func() {
		err = newBuggerDB(connection)
	})

	return err
}

func newBuggerDB(connection string) error {
	var pool *pgxpool.Pool

	var err error

	for i := 0; i < retries; i++ {
		ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		pool, err = pgxpool.Connect(ctxTimeout, connection)
		if err != nil {
			// Logging out each time in case the error changes in between attempts
			logrus.Errorf("Db connect errored %v", err)
			time.Sleep(sleepTime * time.Second)

			continue
		}

		break
	}

	if err != nil {
		return errFailedToConnect
	}

	_buggerDB = &buggerDB{
		db: pool,
	}

	return nil
}

func CloseDatabase() {
	if _buggerDB.db != nil {
		_buggerDB.db.Close()
	}
}
