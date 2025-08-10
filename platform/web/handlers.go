package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong! (fixed by the REAL director 😎)",
	})
}

func (s *Server) handleGetProducts(c *gin.Context) {
	products, err := s.getProductsUsecase.Execute(c.Request.Context())
	if err != nil {
		log.Printf("Internal error from database: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in connecting to database"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (s *Server) handleGetUserByID(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "Fetching data for user",
		"user_id": id,
	})
}