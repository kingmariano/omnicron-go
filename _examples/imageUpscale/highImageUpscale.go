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
	imageFile, err := os.Open("images.jpg")
	if err != nil {
		fmt.Printf("Error opening image file: %v\n", err)
		return
	}
	res, err := client.HighImageUpscaleGeneration(context.Background(), omnicron.HighImageUpscaleGenerationModelAndParams{
		Model: omnicron.ClarityUpscaler,
		Parameters: omnicron.HighImageUpscaleGenerationParams{
			Image:  imageFile,
			Prompt: omnicron.Ptr("masterpiece, best quality, highest"),
		},
	})
	if err != nil {
		fmt.Printf("Error making High ImageUpscaleGeneration request: %v\n", err)
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
