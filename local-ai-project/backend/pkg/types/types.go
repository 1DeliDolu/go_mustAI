// backend/pkg/types/types.go
package types

import (
	"mime/multipart"
)

// WikiResult represents a Wikipedia search result
type WikiResult struct {
	PageID         string  `json:"pageId"`
	Title          string  `json:"title"`
	URL            string  `json:"url"`
	Description    string  `json:"description,omitempty"`
	Extract        string  `json:"extract,omitempty"`
	Thumbnail      string  `json:"thumbnail,omitempty"`
	RelevanceScore float64 `json:"relevanceScore,omitempty"`
}

// Document represents an uploaded document
type Document struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Size       int64  `json:"size"`
	UploadDate string `json:"uploadDate"`
	Status     string `json:"status"`
	Chunks     int    `json:"chunks,omitempty"`
	Embeddings bool   `json:"embeddings,omitempty"`
}

// Model represents an AI model
type Model struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Size             string  `json:"size"`
	Status           string  `json:"status"`
	DownloadProgress float64 `json:"downloadProgress,omitempty"`
	Description      string  `json:"description,omitempty"`
	ModelType        string  `json:"modelType"`
}

// QueryRequest represents a query request
type QueryRequest struct {
	Query            string `json:"query"`
	ModelName        string `json:"model_name"`
	IncludeWiki      bool   `json:"include_wiki"`
	IncludeDocuments bool   `json:"include_documents"`
	MaxSources       int    `json:"max_sources,omitempty"`
}

// QueryResponse represents a query response
type QueryResponse struct {
	Response string `json:"response"`
	Sources  struct {
		Documents []Document   `json:"documents"`
		Wiki      []WikiResult `json:"wiki"`
	} `json:"sources"`
	ModelUsed      string  `json:"modelUsed"`
	ProcessingTime float64 `json:"processingTime"`
}

// Request types
type DownloadModelRequest struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"url" binding:"required"`
}

type LoadModelRequest struct {
	Name string `json:"name" binding:"required"`
}

type UploadDocumentRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// Response types
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
