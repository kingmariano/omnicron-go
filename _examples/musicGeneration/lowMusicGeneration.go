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
	res, err := client.LowMusicGeneration(context.Background(), omnicron.LowMusicGenerationModelAndParams{
		Model: omnicron.RiffusionModel,
		Parameters: &omnicron.LowMusicGenerationParams{
			PromptA: "funky synth solo bass",
			Alpha: omnicron.Ptr(0.6),
		},
	})
	if err!= nil {
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
