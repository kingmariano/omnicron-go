package omnicron

import "errors"

var NoQueryParamter = errors.New("No query parameter specified")
var ModelNotFoundError = errors.New("Model not found")
var GroqChatCompletionNoMessageError = errors.New("GroqChatCompletionError: Message field is required")

type ErrorResponse struct {
	Error string `json:"error"`
}