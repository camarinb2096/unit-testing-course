package server

import (
	"camarinb2096/unit-testing-course/internal/adapter"
	"log"
	"net/http"

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
	api := s.router.Group("/api/v1")
	api.GET("/pokemon/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := adapter.GetPokemon(id)
		if err != nil {
			log.Println("Error:", err)
			c.JSON(500, gin.H{
				"error": "Internal server error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Pokemon data retrieved successfully!"})
	})
}

func (s *Server) Start() {
	log.Println("Server is running...")
	s.router.Run(":8080")
}
