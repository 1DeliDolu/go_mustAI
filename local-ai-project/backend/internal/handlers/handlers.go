// backend/internal/handlers/handlers.go
package handlers

import (
	"net/http"
	"strconv"
	"time"

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
		"status":    "ok",
		"timestamp": time.Now().Unix(),
		"message":   "Local AI Project API is running",
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
	var req struct {
		Name string `json:"name" binding:"required"`
		URL  string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.modelService.DownloadModel(req.Name, req.URL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Model downloaded successfully"})
}

func (h *Handler) LoadModel(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.modelService.LoadModel(req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Model loaded successfully"})
}

func (h *Handler) DeleteModel(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Model name is required"})
		return
	}

	if err := h.modelService.DeleteModel(name); err != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"message":  "Document uploaded successfully",
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

	if err := h.documentService.DeleteDocument(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}

// Wiki handlers
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

	startTime := time.Now()

	// Search documents if requested
	var documents []types.Document
	if req.IncludeDocuments {
		docs, err := h.documentService.SearchDocuments(req.Query)
		if err == nil {
			documents = docs
		}
	}

	// Search wiki if requested
	var wikiResults []types.WikiResult
	if req.IncludeWiki {
		wiki, err := h.wikiService.Search(req.Query)
		if err == nil {
			wikiResults = wiki
		}
	}

	// Generate AI response
	response, err := h.aiService.GenerateResponse(req.Query, documents, wikiResults)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	processingTime := time.Since(startTime).Seconds()

	result := types.QueryResponse{
		Response:       response,
		ModelUsed:      h.aiService.GetCurrentModel(),
		ProcessingTime: processingTime,
	}
	result.Sources.Documents = documents
	result.Sources.Wiki = wikiResults

	c.JSON(http.StatusOK, result)
}
