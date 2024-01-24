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
}

func (s *Server) Run() {
	s.InitRoute()
	if err := s.engine.Run(":8080"); err != nil {
		panic(fmt.Errorf("Failed to start server: %v", err))
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(fmt.Errorf("config error: %v", err))
	}
	db := config.ConnectDB()
	var roomRepository repository.RoomRepository 
	var roomUC usecase.RoomUseCase               
	roomRepository = repository.NewRoomRepository(db)
	roomUC = usecase.NewRoomUseCase(roomRepository)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		roomUC: roomUC,
		engine: engine,
		host:   host,
	}
}
