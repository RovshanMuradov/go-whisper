package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rovshanmuradov/go-whisper-ai/api/whisper"
)

func main() {
	client := whisper.NewClient(whisper.WithKey(os.Getenv("OPENAI_API_KEY")))

	response, err := client.TranscribeFile("file.m4a")
	if err != nil {
		log.Fatalf("Error transcribing file: %v", err)
	}

	fmt.Printf("Transcription: %s\n", response.Text)
}
