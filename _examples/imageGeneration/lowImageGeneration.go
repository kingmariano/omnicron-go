package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/kingmariano/omnicron-go"
)

func main() {
	apiKey := "Your-API-Key"
	client := omnicron.NewClient(apiKey, omnicron.WithBaseURL("https://omnicron-latest.onrender.com/"))
	res, err := client.LowImageGeneration(context.Background(), omnicron.LowImageGenerationModelAndParams{
		Model: omnicron.SDXLLightning4stepModel,
		Parameters: &omnicron.LowImageGenerationParams{
			Prompt:            "self-portrait of a woman, lightning in the background",
	
		},
	})
	if err != nil {
		fmt.Printf("Error making Low ImageGeneration request: %v\n", err)
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
	// dynamically handle the response with the Gabs library: https://github.com/Jeffail/gabs/
	outputText := res.Path("response.output").Data().(map[string]interface{})
	fmt.Println(outputText)

}
