package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/kingmariano/omnicron-go"
)

func main() {
	apiKey := "YOUR_API_KEY"
	client := omnicron.NewClient(apiKey, omnicron.WithBaseURL("https://omnicron-latest.onrender.com/"))
	musicFile, err := os.Open("sample1.wav") // smaple music file
	if err != nil {
		fmt.Printf("Error opening image file: %v\n", err)
		return
	}
	res, err := client.HighMusicGeneration(context.Background(), omnicron.HighMusicGenerationModelAndParams{
		Model: omnicron.MetaMusicGenModel,
		Parameters: omnicron.HighMusicGenerationParams{
			Prompt:         "Create a vibrant Afrobeat track inspired by Nigerian music culture ",
			TopK:           omnicron.Ptr(255),
			ModelVersion:   omnicron.Ptr("melody-large"),
			InputAudioFile: musicFile,
		},
	})
	if err != nil {
		fmt.Printf("Error making HighImageGeneration request: %v\n", err)
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
