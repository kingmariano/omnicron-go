package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/kingmariano/omnicron-go"
)

func main() {
	apiKey := "HwikyHTrh4DMP9ZV1PzOn3K+C4Il7N/RKurD5AjyoIE="
	client := omnicron.NewClient(apiKey, omnicron.WithBaseURL("https://omnicron-latest.onrender.com/"))

	res, err := client.VideoGeneration(context.Background(), omnicron.HighVideoGenerationModelAndParams{
		Model: omnicron.ZeroScopeV2XLModel,
		Parameters: omnicron.HighVideoGenerationParams{
			Prompt: "Create a 30-second animated video showcasing a futuristic cityscape at night, with flying cars, neon lights, and bustling streets.",
		},
	})
	if err != nil {
		fmt.Printf("Error making Video Generation request: %v\n", err)
		return
	}

	// Marshal the response to JSON
	jsonData, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	// Write JSON data to a file
	file, err := os.Create("file.json")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Error writing JSON to file: %v\n", err)
		return
	}

	fmt.Println("JSON data written to file.json successfully")
}
