package models

import (
	"errors"
	"time"
)

type CreateTrackRequest struct {
	Urge       int       `json:"urge"`
	Need       int       `json:"need"`
	CreateTime time.Time `json:"create_time"`
}

func (req *CreateTrackRequest) Valid() error {
	if req.Urge == 0 {
		return errors.New("need to specify an urge")
	}

	if req.Need == 0 {
		return errors.New("need to specify a need")
	}

	if req.CreateTime.IsZero() {
		return errors.New("need to specify a time")
	}

	return nil
}
