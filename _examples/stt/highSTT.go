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
	audioFile, err := os.Open("sample.mp3")
	if err != nil {
		fmt.Printf("Error opening image file: %v\n", err)
		return
	}
	res, err := client.HighSTTGeneration(context.Background(), omnicron.HighSTTModelAndParams{
		Model: omnicron.InsanelyFastWhisperWithVideoModel,
		Parameters: omnicron.HighSTTParams{
			AudioFile: audioFile,
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
	// dynamically handle the response with the Gabs library: https://github.com/Jeffail/gabs/
	outputText := res.Path("response.output").Data().(map[string]interface{})
	fmt.Println(outputText)

}
