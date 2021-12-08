package app

import (
	"github.com/gin-gonic/gin"
	"github.com/parag08/go-errors/pkg/api/sandwich"
	"github.com/parag08/go-errors/pkg/logger"
)

type Server struct {
	router                 *gin.Engine
	sanwdwitchMakerService sandwich.SandwichService
}

func NewServer(router *gin.Engine, sandwichService sandwich.SandwichService) *Server {
	return &Server{
		router:                 router,
		sanwdwitchMakerService: sandwichService,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	err := r.Run()

	if err != nil {
		logger.Log.Error("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
