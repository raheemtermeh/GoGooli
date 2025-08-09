package web

import (
	"github.com/gin-gonic/gin"
	"github.com/raheemtermeh/goyan-framework/internal/usecase/product"
)

type Server struct {
	router           *gin.Engine
	getProductsUsecase *product.GetProductsUsecase
}

func NewServer(getProductsUC *product.GetProductsUsecase) *Server {
	server := &Server{
		getProductsUsecase: getProductsUC,
	}

	router := gin.Default()
	server.setupRoutes(router)
	server.router = router
	return server
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) setupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/ping", s.handlePing)
		api.GET("/products", s.handleGetProducts)
	}
}