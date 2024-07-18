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
	imageFile, err := os.Open("image.jpg") //example image, edit
	if err!= nil {
        fmt.Printf("Error opening image file: %v\n", err)
        return
    }
	res, err := client.HighImageGeneration(context.Background(), omnicron.HighImageGenerationModelAndParams{
		Model: omnicron.PlaygroundV251024pxAestheticModel,
		Parameters: omnicron.HighImageGenerationParams{
			Prompt: "Astronaut in a jungle, cold color palette, muted colors, detailed, 8k",
			NegativePrompt: omnicron.Ptr("ugly, deformed, noisy, blurry, distorted"),
			ImageFile: imageFile,
           Width: omnicron.Ptr(1024),
		   Height: omnicron.Ptr(1024),
		   NumOutputs: omnicron.Ptr(1),
		   Scheduler:         omnicron.Ptr("DPMSolver++"),
           NumInferenceSteps: omnicron.Ptr(30),
		   GuidanceScale: omnicron.Ptr(3.1),
		   PromptStrength: omnicron.Ptr(0.7),
		},
	})
	if err!= nil {
        fmt.Printf("Error making High ImageGeneration request: %v\n", err)
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
