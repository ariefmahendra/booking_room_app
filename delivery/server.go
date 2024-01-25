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
	facilitiesUC       usecase.FacilitiesUsecase
	trxRsvpUC          usecase.TrxRsvUsecase
	reportUC           usecase.ReportUsecase
	engine             *gin.Engine
	host               string
}

func (s *Server) InitRoute() {
	rsvp := s.engine.Group("/rsvp")
	controller.NewTrxRsvpController(s.trxRsvpUC, rsvp).Route()

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

	rp := s.engine.Group("/api/v1/report")
	controller.NewReportController(s.reportUC, rp).Route()

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
	trxRsvpRepo := repository.NewTrxRsvRepository(db)
	reportRepo := repository.NewReportRepository(db)

	employeeUC := usecase.NewEmployeeUC(employeeRepository)
	facilitiesUC := usecase.NewFacilitiesUsecase(facilitiesRepository)
	trxRsvpUC := usecase.NewTrxRsvUseCase(trxRsvpRepo)
	reportUC := usecase.NewReportUsecase(reportRepo)

	employeeController := controller.NewEmployeeController(employeeUC)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		employeeController: employeeController,
		facilitiesUC:       facilitiesUC,
		trxRsvpUC:          trxRsvpUC,
		reportUC:           reportUC,
		engine:             engine,
		host:               host,
	}
}
