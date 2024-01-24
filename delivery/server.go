package delivery

import (
	"booking-room/config"
	"booking-room/delivery/controller"
	"booking-room/repository"
	"booking-room/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	employeeController *controller.EmployeeControllerImpl
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
	employeeUC := usecase.NewEmployeeUC(employeeRepository)
	employeeController := controller.NewEmployeeController(employeeUC)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		employeeController: employeeController,
		engine:             engine,
		host:               host,
	}
}
