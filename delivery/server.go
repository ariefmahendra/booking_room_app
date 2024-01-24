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
	trxRsvpUC  usecase.TrxRsvUsecase
	engine *gin.Engine
	host string
}


func (s *Server) InitRoute(){
	rg := s.engine.Group("/rsvp")
	controller.NewTrxRsvpController(s.trxRsvpUC, rg).Route()
}

func (s *Server) Run() {
	s.InitRoute()
	if err := s.engine.Run(":8080"); err != nil {
		panic(fmt.Errorf("Failed to start server: %v", err))
	}
}

func NewServer() *Server{
	cfg, err := config.NewConfig()
	if err !=nil{
		panic(fmt.Errorf("config error : %v", err))
	}
	db := config.ConnectDB()
	// Inject DB ke -> Repository
		trxRsvpRepo := repository.NewTrxRsvRepository(db)
	
	// Inject Repository ke -> Usecase
		trxRsvpUC := usecase.NewTrxRsvUseCase(trxRsvpRepo)

	// ROUTE
		engine := gin.Default()
		host := fmt.Sprintf(":%s",cfg.ApiPort )

	return &Server{
		trxRsvpUC : trxRsvpUC,
		engine: engine,
		host: host,
	}
}