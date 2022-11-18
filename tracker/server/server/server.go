package server

import (
	"log"
	"tracker/service"

	"github.com/gin-gonic/gin"
)

type TrackerServer interface {
	Run()
}

type trackerServer struct {
	srv service.TrackerService
	eng *gin.Engine
}

func NewTracker(srv service.TrackerService) TrackerServer {
	r := gin.Default()

	log.Println("Created the server")

	return &trackerServer{
		srv: srv,
		eng: r,
	}
}

func (t *trackerServer) Run() {
	t.eng.Run()
}
