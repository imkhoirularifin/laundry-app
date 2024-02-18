package delivery

import (
	"fmt"
	"laundry-app/config"
	"laundry-app/delivery/controller"
	"laundry-app/manager"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	ucManager manager.UsecaseManager
	engine    *gin.Engine
	host      string
}

func (s *Server) setupController() {
	rg := s.engine.Group("/api")

	// register all controller below
	controller.NewEmployeeController(s.ucManager.EmployeeUsecase(), rg).Route()
}

/*
	Method Receiver syntax
*/
// func (s *Server) Start() {
// 	s.setupController()
// 	if err := s.engine.Run(s.host); err != nil {
// 		panic(err)
// 	}
// }

func Start(s *Server) {
	// Start the controller
	s.setupController()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	// Validate environment configuration
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// Validate database config & open connection
	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		panic(err)
	}

	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	// Overwrite Original Server
	return &Server{
		ucManager: usecaseManager,
		engine:    engine,
		host:      host,
	}
}
