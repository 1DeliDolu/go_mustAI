// backend/internal/config/config.go
package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	Port         string
	ModelsPath   string
	UploadsPath  string
	DatabasePath string
	OllamaURL    string
	MaxFileSize  int64
	AllowedTypes []string
}

func Load() *Config {
	// Get port from environment or default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	// Default paths
	homeDir, _ := os.UserHomeDir()
	appDir := filepath.Join(homeDir, ".local-ai-project")

	// Get database path from environment or default
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = filepath.Join(appDir, "data", "app.db")
	}

	// Create directories if they don't exist
	os.MkdirAll(filepath.Join(appDir, "models"), 0755)
	os.MkdirAll(filepath.Join(appDir, "uploads"), 0755)
	os.MkdirAll(filepath.Join(appDir, "data"), 0755)

	return &Config{
		Port:         port,
		ModelsPath:   filepath.Join(appDir, "models"),
		UploadsPath:  filepath.Join(appDir, "uploads"),
		DatabasePath: dbPath,
		OllamaURL:    getEnv("OLLAMA_URL", "http://localhost:11434"),
		MaxFileSize:  50 * 1024 * 1024, // 50MB
		AllowedTypes: []string{".pdf", ".txt", ".docx", ".md"},
	}
}

func NewConfig() *Config {
	return Load()
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
