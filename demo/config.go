package main

import "fmt"

// Config yapılandırma ayarlarını tutar
type Config struct {
	ModelPath   string
	Temperature float64
	MaxTokens   int
	Prompt      string
}

// DefaultConfig varsayılan yapılandırma değerlerini döndürür
func DefaultConfig() Config {
	return Config{
		ModelPath:   "./models/llama-2-7b.Q4_0.bin",
		Temperature: 0.7,
		MaxTokens:   128,
		Prompt:      "Türkiye'nin başkenti neresidir?",
	}
}

// ValidateConfig yapılandırma değerlerini doğrular
func (c *Config) ValidateConfig() error {
	if c.ModelPath == "" {
		return fmt.Errorf("model yolu boş olamaz")
	}
	if c.Temperature < 0 || c.Temperature > 2 {
		return fmt.Errorf("temperature değeri 0-2 arasında olmalıdır")
	}
if c.MaxTokens <= 0 {
	return fmt.Errorf("max tokens değeri 0'dan büyük olmalıdır")
}
return nil
}