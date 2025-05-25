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

func NewModelService(cfg *config.Config, db *sql.DB) *ModelService {
	return &ModelService{config: cfg, db: db}
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
				ID:        file.Name(),
				Name:      file.Name(),
				Size:      fmt.Sprintf("%d MB", info.Size()/(1024*1024)),
				Status:    "available",
				ModelType: "chat",
			})
		}
	}

	return models, nil
}

func (s *ModelService) DownloadModel(name, url string) error {
	// Create the models directory if it doesn't exist
	if err := os.MkdirAll(s.config.ModelsPath, 0755); err != nil {
		return fmt.Errorf("failed to create models directory: %w", err)
	}

	// Download the model file
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download model: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download model: HTTP %d", resp.StatusCode)
	}

	// Create the destination file
	filePath := filepath.Join(s.config.ModelsPath, name)
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create model file: %w", err)
	}
	defer out.Close()

	// Copy the response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		// Clean up partial file on error
		os.Remove(filePath)
		return fmt.Errorf("failed to save model file: %w", err)
	}

	return nil
}

func (s *ModelService) LoadModel(name string) error {
	filePath := filepath.Join(s.config.ModelsPath, name)

	// Check if model file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("model %s not found", name)
	}

	// TODO: Implement actual model loading logic
	// For now, just return success
	return nil
}

func (s *ModelService) DeleteModel(name string) error {
	filePath := filepath.Join(s.config.ModelsPath, name)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("model %s not found", name)
	}

	// Delete the model file
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete model %s: %w", name, err)
	}
	return nil
}
