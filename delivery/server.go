package delivery

import (
	"booking-room/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	host string
}


func (s *Server) InitRoute(){
	rg := s.engine.Group("/")
	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
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

func NewServer() *Server{
	cfg, err := config.NewConfig()
	if err !=nil{
		panic(fmt.Errorf("config error : %v", err))
	}
	config.ConnectDB()
	// Inject DB ke -> Repository

	
	// Inject Repository ke -> Usecase


	// ROUTE
		engine := gin.Default()
		host := fmt.Sprintf(":%s",cfg.ApiPort )

	return &Server{
		engine: engine,
		host: host,
	}
}