package models

import "time"

type CreateTrackRequest struct {
	Urge       int       `json:"urge"`
	Need       int       `json:"need"`
	CreateTime time.Time `json:"create_time"`
}

type CreateTrackResponse struct {
	Message string `json:"message"`
}

type CreateTrackErrorResponse struct {
	Error string `json:"error"`
}
