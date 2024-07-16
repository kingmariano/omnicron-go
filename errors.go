package omnicron

import "errors"

var ErrNoQueryParameter = errors.New("no query parameter specified")
var ErrModelNotFound = errors.New("model not found")
var ErrGroqChatCompletionNoMessage = errors.New("message field is required")
var ErrPromptMissing = errors.New("prompt field is required")

// ErrorResponse represents an error response from the API.
type ErrorResponse struct {
	Error string `json:"error"`
}
