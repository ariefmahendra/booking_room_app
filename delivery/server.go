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
	roomUC usecase.RoomUseCase
	employeeController *controller.EmployeeControllerImpl
	facilitiesUC       usecase.FacilitiesUsecase
	roomController *controller.RoomController
	trxRsvpUC  usecase.TrxRsvUsecase
	engine *gin.Engine
	host string
}


func (s *Server) InitRoute(){
	// route for management employee
	er := s.engine.Group("/api/v1/employees")
	er.POST("/", s.employeeController.CreateEmployee)
	er.PATCH("/:id", s.employeeController.UpdateEmployee)
	er.DELETE("/:id", s.employeeController.DeleteEmployee)
	er.GET("/:id", s.employeeController.GetEmployeeById)
	er.GET("/email/:email", s.employeeController.GetEmployeeByEmail)
	er.GET("/", s.employeeController.GetEmployees)

	// route for management room
	rg := s.engine.Group("/api/v1//room")
	controller.NewRoomController(s.roomUC, rg).Route()
	rg.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// route for management facilities
	fg := s.engine.Group("/api/v1/facilities")
	controller.NewFacilitiesController(s.facilitiesUC, fg).Route()

	// route for management transaction
	rs := s.engine.Group("/api/v1//reservation")
	controller.NewTrxRsvpController(s.trxRsvpUC, rs).Route()
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
	// Inject DB ke -> Repository
	var roomRepository repository.RoomRepository
	roomRepository = repository.NewRoomRepository(db)
	facilitiesRepository := repository.NewFacilitiesRepository(db) 
	employeeRepository := repository.NewEmployeeRepository(db)
	trxRsvpRepo := repository.NewTrxRsvRepository(db)
	
	// Inject Repository ke -> Usecase
	var roomUC usecase.RoomUseCase    
	faciltiiesUC := usecase.NewFacilitiesUsecase(facilitiesRepository)
	employeeUC := usecase.NewEmployeeUC(employeeRepository)
	roomUC = usecase.NewRoomUseCase(roomRepository)	  
	trxRsvpUC := usecase.NewTrxRsvUseCase(trxRsvpRepo)

	employeeController := controller.NewEmployeeController(employeeUC)

	// ROUTE
		engine := gin.Default()
		host := fmt.Sprintf(":%s",cfg.ApiPort )

	return &Server{
		roomUC: roomUC,
		employeeController: employeeController,
		facilitiesUC:       faciltiiesUC,
		trxRsvpUC : trxRsvpUC,
		engine: engine,
		host: host,
	}
}
