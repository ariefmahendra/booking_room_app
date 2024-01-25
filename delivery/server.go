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
<<<<<<< HEAD
	roomController *controller.RoomController
	engine *gin.Engine
	host   string
	roomUC usecase.RoomUseCase
}

func (s *Server) InitRoute() {
	rg := s.engine.Group("/room")
	controller.NewRoomController(s.roomUC, rg).Route()
	rg.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
=======
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
>>>>>>> ed7e6ada7c231957f8498b03fd926752a5f88f1d
}

func (s *Server) Run() {
	s.InitRoute()
<<<<<<< HEAD
	if err := s.engine.Run(":8080"); err != nil {
		panic(fmt.Errorf("Failed to start server: %v", err))
=======
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
>>>>>>> ed7e6ada7c231957f8498b03fd926752a5f88f1d
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
<<<<<<< HEAD
		panic(fmt.Errorf("config error: %v", err))
	}
	db := config.ConnectDB()
	var roomRepository repository.RoomRepository 
	var roomUC usecase.RoomUseCase               
	roomRepository = repository.NewRoomRepository(db)
	roomUC = usecase.NewRoomUseCase(roomRepository)
=======
		panic(fmt.Errorf("config error : %v", err))
	}
	db := config.ConnectDB()

	employeeRepository := repository.NewEmployeeRepository(db)
	facilitiesRepository := repository.NewFacilitiesRepository(db)
	employeeUC := usecase.NewEmployeeUC(employeeRepository)
	faciltiiesUC := usecase.NewFacilitiesUsecase(facilitiesRepository)
	employeeController := controller.NewEmployeeController(employeeUC)
>>>>>>> ed7e6ada7c231957f8498b03fd926752a5f88f1d

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
<<<<<<< HEAD
		roomUC: roomUC,
		engine: engine,
		host:   host,
=======
		employeeController: employeeController,
		facilitiesUC:       faciltiiesUC,
		engine:             engine,
		host:               host,
>>>>>>> ed7e6ada7c231957f8498b03fd926752a5f88f1d
	}
}
