package services

import (
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"local-ai-project/backend/internal/config"
	"local-ai-project/backend/pkg/types"
)

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
		var createdAt string
		err := rows.Scan(&doc.ID, &doc.Name, &doc.Name, &doc.Size, &doc.Type, &createdAt)
		if err != nil {
			return nil, err
		}
		doc.UploadDate = createdAt
		doc.Status = "ready"
		documents = append(documents, doc)
	}

	return documents, nil
}

func (s *DocumentService) UploadDocument(fileHeader *multipart.FileHeader) (*types.Document, error) {
	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll(s.config.UploadsPath, 0755); err != nil {
		return nil, err
	}

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
		ID:         int(id),
		Name:       fileHeader.Filename,
		Type:       filepath.Ext(fileHeader.Filename),
		Size:       fileHeader.Size,
		UploadDate: time.Now().Format("2006-01-02 15:04:05"),
		Status:     "ready",
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
		var filename, content, createdAt string
		err := rows.Scan(&doc.ID, &filename, &doc.Name, &doc.Size,
			&doc.Type, &content, &createdAt)
		if err != nil {
			return nil, err
		}
		doc.UploadDate = createdAt
		doc.Status = "ready"
		documents = append(documents, doc)
	}

	return documents, nil
}

func (s *DocumentService) DeleteDocument(id int) error {
	// Start a transaction for atomic operation
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	// Get document info first
	var filePath string
	err = tx.QueryRow("SELECT path FROM documents WHERE id = ?", id).Scan(&filePath)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("document with id %d not found", id)
		}
		return fmt.Errorf("failed to get document path: %w", err)
	}

	// Delete from database first
	result, err := tx.Exec("DELETE FROM documents WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete document from database: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("document with id %d not found", id)
	}

	// Commit the database transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Delete file from filesystem (after successful database deletion)
	if filePath != "" {
		if err := os.Remove(filePath); err != nil {
			// Log the error but don't fail the operation
			// since the database record is already deleted
			fmt.Printf("Warning: failed to delete file %s: %v\n", filePath, err)
		}
	}

	return nil
}
