package config

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/yaminmuhammad/pay-app/shared/service"
	"github.com/yaminmuhammad/pay-app/usecase"
)

type Server struct {
	customerUC usecase.CustomerUseCase
	jwtService service.JwtService
	engine     *gin.Engine
	port       string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(Config{})
}
