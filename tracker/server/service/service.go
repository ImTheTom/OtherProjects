package service

import (
	"log"
	"tracker/db"
	"tracker/models"
)

type TrackerService interface {
	CreateTrack(req *models.CreateTrackRequest) error
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

func (s *trackerService) CreateTrack(req *models.CreateTrackRequest) error {
	log.Println("creating track")

	if err := s.dbRepo.CreateTrack(req); err != nil {
		return err
	}

	return nil
}
