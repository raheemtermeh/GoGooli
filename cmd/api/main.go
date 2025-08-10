package main

import (
	"context"
	"log"

	"github.com/raheemtermeh/goyan-framework/internal/config" 
	"github.com/raheemtermeh/goyan-framework/internal/usecase/product"
	"github.com/raheemtermeh/goyan-framework/platform/database"
	"github.com/raheemtermeh/goyan-framework/platform/database/sqlc_generated"
	"github.com/raheemtermeh/goyan-framework/platform/web"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	dbpool, err := database.NewDBPool(context.Background(), cfg.Database.URL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()
	log.Println("Successfully connected to the database!")

	queries := sqlc_generated.New(dbpool)
	productRepo := database.NewProductRepository(queries)
	getProductsUsecase := product.NewGetProductsUsecase(productRepo)

	server := web.NewServer(getProductsUsecase, cfg)
	log.Printf("Server is starting on %s...", cfg.Server.Port)

	if err := server.Start(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}