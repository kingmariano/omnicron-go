package omnicron

import (
	"context"
	"encoding/json"
)

type ToolChoice string

const (
	ToolChoiceAuto ToolChoice = "auto"
	ToolChoiceNone ToolChoice = "none"
)

type Message struct {
	Content    string            `json:"content"` // Required fields, not omitting in JSON
	Role       string            `json:"role"`    // Required fields, not omitting in JSON
	Name       string            `json:"name,omitempty"`
	ToolCallID string            `json:"tool_call_id,omitempty"`
	ToolCalls  []MessageToolCall `json:"tool_calls,omitempty"`
}
type MessageToolCall struct {
	ID       string                  `json:"id,omitempty"`
	Function MessageToolCallFunction `json:"function,omitempty"`
	Type     string                  `json:"type,omitempty"`
}

type MessageToolCallFunction struct {
	Arguments string `json:"arguments,omitempty"`
	Name      string `json:"name,omitempty"`
}

type ResponseFormat struct {
	Type string `json:"type,omitempty"`
}

type ToolChoiceToolChoice struct {
	Function ToolChoiceToolChoiceFunction `json:"function,omitempty"`
	Type     string                       `json:"type,omitempty"`
}

type ToolChoiceToolChoiceFunction struct {
	Name string `json:"name,omitempty"`
}

type Tool struct {
	Function ToolFunction `json:"function,omitempty"`
	Type     string       `json:"type,omitempty"`
}

type ToolFunction struct {
	Description string                 `json:"description,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Parameters  map[string]interface{} `json:"parameters,omitempty"`
}

// CompletionCreateParams represents the inputs for the groq completion params.
type CompletionCreateParams struct {
	Messages         []Message      `json:"messages"`
	Model            string         `json:"model"`
	FrequencyPenalty float32        `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	Logprobs         bool           `json:"logprobs,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	N                int            `json:"n,omitempty"`
	PresencePenalty  float32        `json:"presence_penalty,omitempty"`
	ResponseFormat   ResponseFormat `json:"response_format,omitempty"`
	Seed             int            `json:"seed,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	Temperature      float32        `json:"temperature,omitempty"`
	ToolChoice       ToolChoice     `json:"tool_choice,omitempty"`
	Tools            []Tool         `json:"tools,omitempty"`
	TopLogprobs      int            `json:"top_logprobs,omitempty"`
	TopP             float32        `json:"top_p,omitempty"`
	User             string         `json:"user,omitempty"`
}
type GroqChatResponse struct {
	Content string `json:"content"`
}

func (c *Client) GroqChatCompletion(ctx context.Context, req *CompletionCreateParams) (*GroqChatResponse, error) {
	if len(req.Messages) == 0 {
		return nil, GroqChatCompletionNoMessageError
	}
	if req.Model == "" {
		return nil, ModelNotFoundError
	}
	body, err := c.newJSONPostRequest(ctx, "/grok/chatcompletion", "", req)
	if err != nil {
		return nil, err
	}
	var groqChatResponse GroqChatResponse
	if err := json.Unmarshal(body, &groqChatResponse); err != nil {
		return nil, err
	}
	return &groqChatResponse, nil
}
