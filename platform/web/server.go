package web

import (
	"github.com/gin-gonic/gin"
	"github.com/raheemtermeh/goyan-framework/internal/config"
	"github.com/raheemtermeh/goyan-framework/internal/usecase/product"
)

type Server struct {
	router           *gin.Engine
	getProductsUsecase *product.GetProductsUsecase
	cfg                config.Config 

}

func NewServer(getProductsUC *product.GetProductsUsecase, cfg config.Config) *Server {
	server := &Server{
		getProductsUsecase: getProductsUC,
		cfg:                cfg,
	}

	router := gin.Default()
	router.Use(LoggingMiddleware())
	server.setupRoutes(router)
	server.router = router
	return server
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) setupMiddlewares(router *gin.Engine) {

	router.Use(CORSMiddleware())
	router.Use(GzipMiddleware())
	router.Use(LoggingMiddleware())
}

func (s *Server) setupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/ping", s.handlePing)
		api.GET("/products", s.handleGetProducts)
	}

    protected := router.Group("/api/protected")
	protected.Use(AuthMiddleware(s.cfg.Security.JWTSecret)) 
	{
		//test route for jwt middleware 
		protected.GET("/profile", s.handleProfile)
	}
	
}