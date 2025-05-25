// backend/cmd/server/main.go
package main

import (
	"local-ai-project/backend/internal/config"
	"local-ai-project/backend/internal/handlers"
	"local-ai-project/backend/internal/services"
	"local-ai-project/backend/internal/storage"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := storage.InitDB(cfg.DatabasePath)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.Close()
	// Initialize services
	modelService := services.NewModelService(cfg, db)
	documentService := services.NewDocumentService(db, cfg)
	wikiService := services.NewWikiService()
	aiService := services.NewAIService(cfg)

	// Initialize handlers
	h := handlers.New(modelService, documentService, wikiService, aiService)

	// Setup Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Health check
	r.GET("/health", h.HealthCheck)

	// API routes
	api := r.Group("/api/v1")
	{
		// Model management
		models := api.Group("/models")
		{
			models.GET("", h.ListModels)
			models.POST("/download", h.DownloadModel)
			models.POST("/load", h.LoadModel)
			models.DELETE("/:name", h.DeleteModel)
		}

		// Document management
		documents := api.Group("/documents")
		{
			documents.GET("", h.ListDocuments)
			documents.POST("/upload", h.UploadDocument)
			documents.DELETE("/:id", h.DeleteDocument)
		}

		// Wiki search
		wiki := api.Group("/wiki")
		{
			wiki.GET("/search", h.SearchWiki)
		}

		// AI Query
		api.POST("/query", h.Query)
	}

	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
