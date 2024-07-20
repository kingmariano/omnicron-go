package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/kingmariano/omnicron-go"
)

func main() {
	apiKey := "MY_API_KEY"
	client := omnicron.NewClient(apiKey, omnicron.WithBaseURL("https://omnicron-latest.onrender.com/"))

	g4fResponse, err := client.GPT4Free(context.Background(), &omnicron.G4FRequest{
		Messages: []omnicron.Message{
			{
				Content: "What is the weather like in New York City?",
				Role:    "user",
			},
		},
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	jsonData, err := json.MarshalIndent(g4fResponse, "", "  ")
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
