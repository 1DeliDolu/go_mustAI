// backend/internal/handlers/handlers.go
package handlers

import (
	"net/http"
	"strconv"

	"local-ai-project/backend/internal/services"
	"local-ai-project/backend/pkg/types"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	modelService    *services.ModelService
	documentService *services.DocumentService
	wikiService     *services.WikiService
	aiService       *services.AIService
}

func New(modelService *services.ModelService, documentService *services.DocumentService, 
	wikiService *services.WikiService, aiService *services.AIService) *Handler {
	return &Handler{
		modelService:    modelService,
		documentService: documentService,
		wikiService:     wikiService,
		aiService:       aiService,
	}
}

// Health check
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"message": "Local AI API is running",
	})
}

// Model handlers
func (h *Handler) ListModels(c *gin.Context) {
	models, err := h.modelService.ListModels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"models": models})
}

func (h *Handler) DownloadModel(c *gin.Context) {
	var req types.DownloadModelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.modelService.DownloadModel(req.Name, req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Model download started"})
}

func (h *Handler) LoadModel(c *gin.Context) {
	var req types.LoadModelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.aiService.LoadModel(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Model loaded successfully"})
}

func (h *Handler) DeleteModel(c *gin.Context) {
	name := c.Param("name")
	err := h.modelService.DeleteModel(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Model deleted successfully"})
}

// Document handlers
func (h *Handler) ListDocuments(c *gin.Context) {
	documents, err := h.documentService.ListDocuments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"documents": documents})
}

func (h *Handler) UploadDocument(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	document, err := h.documentService.UploadDocument(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Document uploaded successfully",
		"document": document,
	})
}

func (h *Handler) DeleteDocument(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}

	err = h.documentService.DeleteDocument(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}

// Wiki search handler
func (h *Handler) SearchWiki(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	results, err := h.wikiService.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

// AI Query handler
func (h *Handler) Query(c *gin.Context) {
	var req types.QueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Search relevant documents
	documents, err := h.documentService.SearchDocuments(req.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Search wiki if requested
	var wikiResults []types.WikiResult
	if req.IncludeWiki {
		wikiResults, err = h.wikiService.Search(req.Query)
		if err != nil {
			// Don't fail the entire request if wiki search fails
			wikiResults = []types.WikiResult{}
		}
	}

	// Generate AI response
	response, err := h.aiService.GenerateResponse(req.Query, documents, wikiResults)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": response,
		"sources": gin.H{
			"documents": documents,
			"wiki": wikiResults,
		},
	})
}