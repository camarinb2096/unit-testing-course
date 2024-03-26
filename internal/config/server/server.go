package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) Routes() {
	api := s.router.Group("/api")
	api.GET("/pokemon", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Pokemon!"})
	})
}

func (s *Server) Start() {
	log.Println("Server is running...")
	s.router.Run(":8080")
}
