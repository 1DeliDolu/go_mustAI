// backend/internal/config/config.go
package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	Port           string
	DatabasePath   string
	ModelsPath     string
	UploadsPath    string
	OllamaURL      string
	WikiAPIURL     string
	MaxFileSize    int64
	AllowedFileTypes []string
}

func Load() *Config {
	cfg := &Config{
		Port:           getEnv("PORT", "8080"),
		DatabasePath:   getEnv("DB_PATH", "../data/app.db"),
		ModelsPath:     getEnv("MODELS_PATH", "../models"),
		UploadsPath:    getEnv("UPLOADS_PATH", "../uploads"),
		OllamaURL:      getEnv("OLLAMA_URL", "http://localhost:11434"),
		WikiAPIURL:     getEnv("WIKI_API_URL", "https://en.wikipedia.org/api/rest_v1"),
		MaxFileSize:    10 * 1024 * 1024, // 10MB
		AllowedFileTypes: []string{".pdf", ".txt", ".docx", ".md"},
	}

	// Create directories if they don't exist
	createDirIfNotExists(cfg.ModelsPath)
	createDirIfNotExists(cfg.UploadsPath)
	createDirIfNotExists(filepath.Dir(cfg.DatabasePath))

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func createDirIfNotExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
}