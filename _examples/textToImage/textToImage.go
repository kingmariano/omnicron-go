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
    imageFile, err := os.Open("images.jpg")
	if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return     
    }
	imageToTextResponse, err := client.ImageToText(context.Background(), omnicron.ImageToTextParams{
		File: imageFile ,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
        return     
	}

	jsonData, err := json.MarshalIndent(imageToTextResponse, "", "  ")
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
