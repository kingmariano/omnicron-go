package main

import (
	"context"
	"fmt"

	"github.com/kingmariano/omnicron-go"
)

func main() {
	apiKey := "YOUR_API_KEY"
	client := omnicron.NewClient(apiKey, omnicron.WithBaseURL("https://omnicron-latest.onrender.com/"))
	res, err := client.GroqChatCompletion(context.Background(), &omnicron.GroqChatCompletionParams{
		Messages: []omnicron.Message{
			{
				Role:    "user",
				Content: "Hello, what is AI?",
			},
		},
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Print(res.Content)
}
