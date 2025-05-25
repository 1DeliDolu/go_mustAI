package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"local-ai-project/backend/internal/config"
	"local-ai-project/backend/pkg/types"
)

type AIService struct {
	config       *config.Config
	currentModel string
	client       *http.Client
}

func NewAIService(cfg *config.Config) *AIService {
	return &AIService{
		config: cfg,
		client: &http.Client{},
	}
}

func (s *AIService) LoadModel(modelName string) error {
	// For Ollama, we can pull/load the model
	reqBody := map[string]interface{}{
		"name": modelName,
	}

	jsonBody, _ := json.Marshal(reqBody)

	resp, err := http.Post(s.config.OllamaURL+"/api/pull", "application/json",
		bytes.NewReader(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to connect to Ollama: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to load model: HTTP %d", resp.StatusCode)
	}

	s.currentModel = modelName
	return nil
}

func (s *AIService) GenerateResponse(query string, documents []types.Document, wikiResults []types.WikiResult) (string, error) {
	// Build context from documents and wiki results
	var context strings.Builder

	// Add document context
	if len(documents) > 0 {
		context.WriteString("Relevant documents:\n")
		for _, doc := range documents {
			context.WriteString(fmt.Sprintf("- %s\n", doc.Name))
		}
		context.WriteString("\n")
	}

	// Add wiki context
	if len(wikiResults) > 0 {
		context.WriteString("Wikipedia information:\n")
		for _, wiki := range wikiResults {
			context.WriteString(fmt.Sprintf("- %s: %s\n", wiki.Title, wiki.Extract))
		}
		context.WriteString("\n")
	}

	// Create the prompt
	prompt := fmt.Sprintf(`Based on the following context, please answer the question.

Context:
%s

Question: %s

Answer:`, context.String(), query)

	// Call Ollama API
	reqBody := map[string]interface{}{
		"model":  s.currentModel,
		"prompt": prompt,
		"stream": false,
	}

	jsonBody, _ := json.Marshal(reqBody)

	resp, err := http.Post(s.config.OllamaURL+"/api/generate", "application/json",
		bytes.NewReader(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to generate response: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("AI service error: HTTP %d", resp.StatusCode)
	}

	var response struct {
		Response string `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return response.Response, nil
}

func (s *AIService) GetCurrentModel() string {
	return s.currentModel
}
