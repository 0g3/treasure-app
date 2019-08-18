package server

import (
	"github.com/0g3/treasure-app/app/api/v1/router"
	"github.com/0g3/treasure-app/app/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	e *gin.Engine
}

func New() (*Server, error) {
	e := gin.Default()
	router.Route(e)

	return &Server{
		e: e,
	}, nil
}

func (s *Server) Run() error {
	return s.e.Run(":" + config.Env().ServerPort)
}
