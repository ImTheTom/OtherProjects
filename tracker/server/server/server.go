package server

import (
	"log"
	"net/http"
	"tracker/models"
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

	server := &trackerServer{
		srv: srv,
		eng: r,
	}

	r.POST("/create_track", server.handleCreateTrack)

	log.Println("Created the server")

	return server
}

func (t *trackerServer) Run() {
	t.eng.Run()
}

func (s *trackerServer) handleCreateTrack(c *gin.Context) {
	var req models.CreateTrackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Valid(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.srv.CreateTrack(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
