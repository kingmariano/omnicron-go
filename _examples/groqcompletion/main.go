package main

import (
	"context"
	"fmt"

	"github.com/kingmariano/omnicron-go"
)

func main(){
    apiKey := "YOUR_API_KEY"
    client := omnicron.NewClient(apiKey, omnicron.WithBaseURL("https://omnicron-latest.onrender.com"))
	client.GroqChatCompletion(context.Background(), &omnicron.CompletionCreateParams{})
}