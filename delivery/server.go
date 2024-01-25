package delivery

import (
	"booking-room/config"
	"booking-room/delivery/controller"
	"booking-room/repository"
	"booking-room/shared/service"
	"booking-room/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	employeeController *controller.EmployeeControllerImpl
	facilitiesUC       usecase.FacilitiesUsecase
	engine             *gin.Engine
	host               string
}

func (s *Server) InitRoute() {
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// route for management employee
	er := s.engine.Group("/api/v1/employees")
	er.POST("/", s.employeeController.CreateEmployee)
	er.PATCH("/:id", s.employeeController.UpdateEmployee)
	er.DELETE("/:id", s.employeeController.DeleteEmployee)
	er.GET("/:id", s.employeeController.GetEmployeeById)
	er.GET("/email/:email", s.employeeController.GetEmployeeByEmail)
	er.GET("/", s.employeeController.GetEmployees)

	// route for management room

	// route for management facilities
	fg := s.engine.Group("/api/v1/facilities")
	controller.NewFacilitiesController(s.facilitiesUC, fg).Route()

	// route for management transaction
}

func (s *Server) Run() {
	s.InitRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(fmt.Errorf("config error : %v", err))
	}
	db := config.ConnectDB()

	employeeRepository := repository.NewEmployeeRepository(db)
	facilitiesRepository := repository.NewFacilitiesRepository(db)

	employeeUC := usecase.NewEmployeeUC(employeeRepository)
	faciltiiesUC := usecase.NewFacilitiesUsecase(facilitiesRepository)

	employeeController := controller.NewEmployeeController(employeeUC)

	service.NewJwtService(cfg.TokenConfig)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		employeeController: employeeController,
		facilitiesUC:       faciltiiesUC,
		engine:             engine,
		host:               host,
	}
}
