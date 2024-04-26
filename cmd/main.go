package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	atg "github.com/johanesalxd/audio-transcribe-go"
)

func main() {
	funcframework.RegisterHTTPFunction("/", atg.AudioTranscribe)

	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
