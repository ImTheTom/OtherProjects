package service

import (
	"log"
	"tracker/db"
)

type TrackerService interface {
}

type trackerService struct {
	dbRepo db.TrackerDB
}

func NewTracker(
	db db.TrackerDB,
) TrackerService {
	log.Println("Created the service")

	return &trackerService{
		dbRepo: db,
	}
}
