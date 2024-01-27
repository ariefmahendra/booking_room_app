package delivery

import (
	"booking-room/config"
	"booking-room/delivery/controller"
	"booking-room/delivery/middleware"
	"booking-room/repository"
	"booking-room/shared/service"
	"booking-room/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	middleware   *middleware.Middleware
	authUC       usecase.AuthUC
	employeeUC   usecase.EmployeeUC
	roomUC       usecase.RoomUseCase
	facilitiesUC usecase.FacilitiesUsecase
	trxRsvpUC    usecase.TrxRsvUsecase
	reportUC     usecase.ReportUsecase
	engine       *gin.Engine
	host         string
}

func (s *Server) InitRoute() {
	s.engine.Use(s.middleware.NewAuth)

	ar := s.engine.Group("/api/v1/auth")
	controller.NewAuthController(s.authUC, ar).Route()

	// route for management employee
	er := s.engine.Group("/api/v1/employees")
	controller.NewEmployeeController(s.employeeUC, s.middleware, er).Route()

	// route for management room
	rg := s.engine.Group("/api/v1/room")
	controller.NewRoomController(s.roomUC, s.middleware, rg).Route()

	// route for management facilities
	fg := s.engine.Group("/api/v1/facilities")
	controller.NewFacilitiesController(s.facilitiesUC, s.middleware, fg).Route()

	// route for report
	rp := s.engine.Group("/api/v1/report")
	controller.NewReportController(s.reportUC, s.middleware, rp).Route()

	// route for management transaction
	rs := s.engine.Group("/api/v1/reservation")
	controller.NewTrxRsvpController(s.trxRsvpUC, s.middleware, rs).Route()
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

	roomRepository := repository.NewRoomRepository(db)
	facilitiesRepository := repository.NewFacilitiesRepository(db)
	employeeRepository := repository.NewEmployeeRepository(db)
	trxRsvpRepo := repository.NewTrxRsvRepository(db)
	reportRepo := repository.NewReportRepository(db)

	faciltiiesUC := usecase.NewFacilitiesUsecase(facilitiesRepository)
	employeeUC := usecase.NewEmployeeUC(employeeRepository)
	roomUC := usecase.NewRoomUseCase(roomRepository)
	reportUC := usecase.NewReportUsecase(reportRepo)
	roomUC = usecase.NewRoomUseCase(roomRepository)
	trxRsvpUC := usecase.NewTrxRsvUseCase(trxRsvpRepo, roomUC)

	jwtService := service.NewJwtService(cfg.TokenConfig)
	authUC := usecase.NewAuthUC(employeeRepository, jwtService)

	newMiddleware := middleware.NewMiddleware(jwtService)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		middleware:   newMiddleware,
		authUC:       authUC,
		roomUC:       roomUC,
		employeeUC:   employeeUC,
		facilitiesUC: faciltiiesUC,
		trxRsvpUC:    trxRsvpUC,
		reportUC:     reportUC,
		engine:       engine,
		host:         host,
	}
}
