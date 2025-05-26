package main

import (
	"fmt"
	"os"

	llama "github.com/go-skynet/go-llama.cpp"
)

func main() {
	modelPath := "./models/7B/ggml-model-q4_0.bin"

	l, err := llama.New(&llama.Options{
		Model:      modelPath,
		NCtx:       512,
		Seed:       -1,
		F16Memory:  true,
		Embeddings: false,
	})
	if err != nil {
		fmt.Println("Fehler beim Laden des Modells:", err)
		os.Exit(1)
	}

	fmt.Println("Modell erfolgreich geladen.")
	_, err = l.Predict("Was ist künstliche Intelligenz?", llama.SetTokenCallback(func(token string) bool {
		fmt.Print(token)
		return true
	}))
	if err != nil {
		fmt.Println("Fehler bei Vorhersage:", err)
	}
}
