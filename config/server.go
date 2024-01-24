package config

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	engine *gin.Engine
	port   string
}

func (s *Server) initRoute() {
	rg := s.engine.Run()
}
