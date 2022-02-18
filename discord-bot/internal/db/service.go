package db

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ImTheTom/OtherProjects/discord-bot/model"
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
	_discDB    *discordDB
)

type DiscordDBInterface interface {
	UpsertUser(ctx context.Context, user model.User) error
	InsertGamble(ctx context.Context, gamble model.Gamble) error
	FindLatestGambleForUser(ctx context.Context, user model.User) (model.Gamble, error)
	IncreasePoints(ctx context.Context, user model.User) error
	SetUserPoints(ctx context.Context, user model.User) error
	FindByUserIDAndGuildID(ctx context.Context, userID, guildID string) (model.User, error)
	FindTopTenPointsForAGuild(ctx context.Context, guildID string) ([]model.User, error)
}

type discordDB struct {
	db *pgxpool.Pool
	mu sync.Mutex
}

func GetDatabaseInterface() DiscordDBInterface {
	return _discDB
}

func NewDiscordDB(connection string) error {
	var err error

	dbConnOnce.Do(func() {
		err = newDiscordDB(connection)
	})

	return err
}

func newDiscordDB(connection string) error {
	var pool *pgxpool.Pool

	var err error

	for i := 0; i < retries; i++ {
		pool, err = pgxpool.Connect(context.Background(), connection)
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

	_discDB = &discordDB{
		db: pool,
		mu: sync.Mutex{},
	}

	return nil
}

func CloseDatabase() {
	if _discDB.db != nil {
		_discDB.db.Close()
	}
}
