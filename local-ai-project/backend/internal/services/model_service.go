// backend/internal/services/model_service.go
package services

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"local-ai-project/backend/internal/config"
	"local-ai-project/backend/pkg/types"
)

type ModelService struct {
	config *config.Config
	db     *sql.DB
}

func NewModelService(cfg *config.Config) *ModelService {
	return &ModelService{config: cfg}
}

func (s *ModelService) ListModels() ([]types.Model, error) {
	// List downloaded models from filesystem
	var models []types.Model
	
	files, err := os.ReadDir(s.config.ModelsPath)
	if err != nil {
		return models, nil // Return empty list if directory doesn't exist
	}

	for _, file := range files {
		if !file.IsDir() {
			info, _ := file.Info()
			models = append(models, types.Model{
				Name: file.Name(),
				Path: filepath.Join(s.config.ModelsPath, file.Name()),
				Size: info.Size(),
				Status: "downloaded",
			})
		}
	}

	return models, nil
}

func (s *ModelService) DownloadModel(name, url string) error {
	// Create the models directory if it doesn't exist
	if err := os.MkdirAll(s.config.ModelsPath, 0755); err != nil {
		return err
	}

	// Download the model file
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the destination file
	filePath := filepath.Join(s.config.ModelsPath, name)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy the response body to the file
	_, err = io.Copy(out, resp.Body)
	return err
}

func (s *ModelService) DeleteModel(name string) error {
	filePath := filepath.Join(s.config.ModelsPath, name)
	return os.Remove(filePath)
}

// backend/internal/services/document_service.go
type DocumentService struct {
	db     *sql.DB
	config *config.Config
}

func NewDocumentService(db *sql.DB, cfg *config.Config) *DocumentService {
	return &DocumentService{db: db, config: cfg}
}

func (s *DocumentService) ListDocuments() ([]types.Document, error) {
	query := `SELECT id, filename, original_name, size, type, created_at FROM documents ORDER BY created_at DESC`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []types.Document
	for rows.Next() {
		var doc types.Document
		err := rows.Scan(&doc.ID, &doc.Filename, &doc.OriginalName, &doc.Size, &doc.Type, &doc.CreatedAt)
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

func (s *DocumentService) UploadDocument(fileHeader *multipart.FileHeader) (*types.Document, error) {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)
	filePath := filepath.Join(s.config.UploadsPath, filename)

	// Create the destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Copy file content
	if _, err = io.Copy(dst, file); err != nil {
		return nil, err
	}

	// Extract text content
	content, err := s.extractTextContent(filePath, fileHeader.Filename)
	if err != nil {
		content = "" // Continue even if text extraction fails
	}

	// Save to database
	query := `INSERT INTO documents (filename, original_name, path, size, type, content) 
			  VALUES (?, ?, ?, ?, ?, ?)`
	
	result, err := s.db.Exec(query, filename, fileHeader.Filename, filePath, fileHeader.Size, 
		filepath.Ext(fileHeader.Filename), content)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	return &types.Document{
		ID:           int(id),
		Filename:     filename,
		OriginalName: fileHeader.Filename,
		Path:         filePath,
		Size:         fileHeader.Size,
		Type:         filepath.Ext(fileHeader.Filename),
		Content:      content,
	}, nil
}

func (s *DocumentService) extractTextContent(filePath, originalName string) (string, error) {
	ext := filepath.Ext(originalName)
	
	switch ext {
	case ".txt", ".md":
		content, err := os.ReadFile(filePath)
		return string(content), err
	case ".pdf":
		// TODO: Implement PDF text extraction
		return "PDF text extraction not implemented yet", nil
	case ".docx":
		// TODO: Implement DOCX text extraction
		return "DOCX text extraction not implemented yet", nil
	default:
		return "", fmt.Errorf("unsupported file type: %s", ext)
	}
}

func (s *DocumentService) SearchDocuments(query string) ([]types.Document, error) {
	// Simple text search in content
	sqlQuery := `SELECT id, filename, original_name, size, type, content, created_at 
				 FROM documents 
				 WHERE content LIKE ? 
				 ORDER BY created_at DESC LIMIT 5`
	
	rows, err := s.db.Query(sqlQuery, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []types.Document
	for rows.Next() {
		var doc types.Document
		err := rows.Scan(&doc.ID, &doc.Filename, &doc.OriginalName, &doc.Size, 
			&doc.Type, &doc.Content, &doc.CreatedAt)
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

func (s *DocumentService) DeleteDocument(id int) error {
	// Get document info first
	var filePath string
	err := s.db.QueryRow("SELECT path FROM documents WHERE id = ?", id).Scan(&filePath)
	if err != nil {
		return err
	}

	// Delete file from filesystem
	if err := os.Remove(filePath); err != nil {
		// Continue even if file deletion fails
	}

	// Delete from database
	_, err = s.db.Exec("DELETE FROM documents WHERE id = ?", id)
	return err
}