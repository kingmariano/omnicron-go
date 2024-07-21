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
    audioFile, err := os.Open("girl of my dreams rodwave.mp3")
	if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return     
    }
	shazamResponse, err := client.Shazam(context.Background(), omnicron.ShazamParams{
		File: audioFile,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
        return     
	}
	jsonData, err := json.MarshalIndent(shazamResponse, "", "  ")
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
