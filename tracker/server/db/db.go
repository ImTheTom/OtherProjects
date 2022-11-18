package db

import (
	"database/sql"
	"log"
	"tracker/models"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "tracker.db"

type TrackerDB interface {
	CreateTrack(req *models.CreateTrackRequest) error
}

type trackerDB struct {
	db *sql.DB
}

func NewTracker() (TrackerDB, error) {
	log.Println("opening db")

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(createTableQuery); err != nil {
		return nil, err
	}

	log.Println("Connected the db")

	return &trackerDB{
		db: db,
	}, nil
}

func (d *trackerDB) CreateTrack(req *models.CreateTrackRequest) error {
	_, err := d.db.Exec(createTrackQuery, req.Urge, req.Need, req.CreateTime)
	return err
}
