// backend/pkg/types/types.go
package types

import (
	"mime/multipart"
	"time"
)

// Model represents an AI model
type Model struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Size      int64     `json:"size"`
	Status    string    `json:"status"` // downloaded, loading, loaded
	CreatedAt time.Time `json:"created_at"`
}

// Document represents an uploaded document
type Document struct {
	ID           int       `json:"id"`
	Filename     string    `json:"filename"`
	OriginalName string    `json:"original_name"`
	Path         string    `json:"path"`
	Size         int64     `json:"size"`
	Type         string    `json:"type"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
}

// WikiResult represents a Wikipedia search result
type WikiResult struct {
	Title       string `json:"title"`
	Extract     string `json:"extract"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Thumbnail   string `json:"thumbnail"`
}

// Request types
type DownloadModelRequest struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"url" binding:"required"`
}

type LoadModelRequest struct {
	Name string `json:"name" binding:"required"`
}

type QueryRequest struct {
	Query       string `json:"query" binding:"required"`
	IncludeWiki bool   `json:"include_wiki"`
	ModelName   string `json:"model_name"`
}

type UploadDocumentRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// Response types
type QueryResponse struct {
	Response string `json:"response"`
	Sources  struct {
		Documents []Document   `json:"documents"`
		Wiki      []WikiResult `json:"wiki"`
	} `json:"sources"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}