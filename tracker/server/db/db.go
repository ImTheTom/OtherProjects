package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "tracker.db"

type TrackerDB interface {
}

type trackerDB struct {
	db *sql.DB
}

func NewTracker() (TrackerDB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(create); err != nil {
		return nil, err
	}

	log.Println("Connected the db")

	return &trackerDB{
		db: db,
	}, nil
}
