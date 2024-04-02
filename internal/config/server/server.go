package server

import (
	"camarinb2096/unit-testing-course/internal/adapter"
	"camarinb2096/unit-testing-course/internal/dto"
	"log"
	"net/http"
	"strconv"

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
		var requestParams dto.RequestParams
		requestParams.Id, _ = c.Params.Get("id")
		requestParams.Page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
		requestParams.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))

		response := adapter.GetPokemonList(requestParams)
		if response.Status != http.StatusOK {
			switch response.Status {
			case 404:
				c.JSON(http.StatusNotFound, gin.H{"message": "Pokemon not found"})
			case 500:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			case 400:
				c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			}
			return
		}
		c.JSON(http.StatusOK, response)
	})
}

func (s *Server) Start() {
	log.Println("Server is running...")
	s.router.Run(":8080")
}
