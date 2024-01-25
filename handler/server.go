package handler

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/yaminmuhammad/pay-app/config"
	"github.com/yaminmuhammad/pay-app/handler/controller"
	"github.com/yaminmuhammad/pay-app/handler/middleware"
	"github.com/yaminmuhammad/pay-app/repository"
	"github.com/yaminmuhammad/pay-app/shared/service"
	"github.com/yaminmuhammad/pay-app/usecase"
)

type Server struct {
	customerUC usecase.CustomerUseCase
	authUC     usecase.AuthUseCase
	jwtService service.JwtService
	engine     *gin.Engine
	port       string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)

	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewAuthController(s.authUC, rg).Route()
	controller.NewCustomerController(s.customerUC, rg, authMiddleware).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.port); err != nil {
		log.Fatalf("server can't running on port '%v', error : %v", s.port, err)
	}
}

func NewServer() *Server {
	config, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)
	db, err := sql.Open(config.Driver, dsn)
	if err != nil {
		panic("connection error")
	}

	customerRepo := repository.NewCustomerRepo(db)
	jwtService := service.NewJwtService(config.TokenConfig)
	customerUC := usecase.NewCustomerUseCase(customerRepo)
	authUC := usecase.NewAuthUseCase(customerUC, jwtService)

	engine := gin.Default()
	port := fmt.Sprintf(":%s", config.ApiPort)
	return &Server{
		customerUC,
		authUC,
		jwtService,
		engine,
		port,
	}
}