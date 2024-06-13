package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var (
	API_KEY string
)

func generateTextToTextResponse(textRequest string) {
	ctx := context.Background()

	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// The Gemini 1.5 models are versatile and work with most use cases
	model := client.GenerativeModel("gemini-1.5-flash")

	res, err := model.GenerateContent(ctx, genai.Text(textRequest))
	if err != nil {
		fmt.Println("There was an error: ", err)
	}

	// Print the response in a readable format
	for i, candidate := range res.Candidates {
		fmt.Printf("Candidate %d: %s\n", i+1, candidate.Content)
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	API_KEY = os.Getenv("API_KEY")
}

func main() {
	generateTextToTextResponse("Write a story about development of Apache Kafka")
}
