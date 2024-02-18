package delivery

import (
	"fmt"
	"laundry-app/config"
	"laundry-app/delivery/controller"
	"laundry-app/delivery/middleware"
	"laundry-app/manager"
	"laundry-app/utils/common"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	ucManager  manager.UsecaseManager
	engine     *gin.Engine
	host       string
	logService common.MyLogger
}

func (s *Server) setupController() {
	// middleware
	s.engine.Use(middleware.NewLogMiddleware(s.logService).LogRequest())

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
	// Create environment configuration instance
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// Create connection to database
	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		panic(err)
	}

	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	// Services
	logService := common.NewMyLogger(cfg.LogConfig)

	// Overwrite Original Server
	return &Server{
		ucManager:  usecaseManager,
		engine:     engine,
		host:       host,
		logService: logService,
	}
}
