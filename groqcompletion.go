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

type GroqChatCompletionResponse struct {
	ID                string   `json:"id"`
	Choices           []Choice `json:"choices"`
	Created           int      `json:"created"`
	Model             string   `json:"model"`
	Object            string   `json:"object"`
	SystemFingerprint string   `json:"systemFingerprint"`
	Usage             Usage    `json:"usage,omitempty"`

	Stream chan *ChatChunkCompletion `json:"-"`
}

type Choice struct {
	FinishReason string         `json:"finishReason"`
	Index        int            `json:"index"`
	Logprobs     ChoiceLogprobs `json:"logprobs"`
	Message      ChoiceMessage  `json:"message"`
}
type ChoiceMessageToolCall struct {
	ID       *string                       `json:"id,omitempty"`
	Function ChoiceMessageToolCallFunction `json:"function,omitempty"`
	Type     *string                       `json:"type,omitempty"`
}
type ChoiceMessageToolCallFunction struct {
	Arguments *string `json:"arguments,omitempty"`
	Name      *string `json:"name,omitempty"`
}
type ChoiceMessage struct {
	Content   string                  `json:"content"`
	Role      string                  `json:"role"`
	ToolCalls []ChoiceMessageToolCall `json:"toolCalls,omitempty"`
}
type ChoiceLogprobs struct {
	Content []ChoiceLogprobsContent `json:"content,omitempty"`
}
type ChoiceLogprobsContent struct {
	Token       *string                           `json:"token,omitempty"`
	Bytes       []int                             `json:"bytes,omitempty"`
	Logprob     *float64                          `json:"logprob,omitempty"`
	TopLogprobs []ChoiceLogprobsContentTopLogprob `json:"topLogprobs,omitempty"`
}
type ChoiceLogprobsContentTopLogprob struct {
	Token   *string  `json:"token,omitempty"`
	Bytes   []int    `json:"bytes,omitempty"`
	Logprob *float64 `json:"logprob,omitempty"`
}

type ChatChunkCompletion struct {
	ID                *string       `json:"id"`
	Choices           []ChoiceChunk `json:"choices"`
	Created           *int          `json:"created"`
	Model             *string       `json:"model"`
	Object            *string       `json:"object"`
	SystemFingerprint *string       `json:"systemFingerprint"`
	XGroq             *XGroq        `json:"xGroq,omitempty"`
	// contains filtered or unexported fields
}
type XGroq struct {
	Usage Usage `json:"usage"`
}
type ChoiceChunk struct {
	Delta        ChoiceDelta    `json:"delta"`
	FinishReason string         `json:"finishReason"`
	Index        int            `json:"index"`
	Logprobs     ChoiceLogprobs `json:"logprobs"`
}
type ChoiceDelta struct {
	Content      string                   `json:"content"`
	Role         string                   `json:"role"`
	FunctionCall *ChoiceDeltaFunctionCall `json:"functionCall,omitempty"`
	ToolCalls    []ChoiceDeltaToolCall    `json:"toolCalls,omitempty"`
}
type ChoiceDeltaFunctionCall struct {
	Arguments *string `json:"arguments,omitempty"`
	Name      *string `json:"name,omitempty"`
}
type ChoiceDeltaToolCall struct {
	Index    int                         `json:"index"`
	ID       *string                     `json:"id,omitempty"`
	Function ChoiceDeltaToolCallFunction `json:"function,omitempty"`
	Type     *string                     `json:"type,omitempty"`
}
type ChoiceDeltaToolCallFunction struct {
	Arguments *string `json:"arguments,omitempty"`
	Name      *string `json:"name,omitempty"`
}
type Usage struct {
	CompletionTime   *float64 `json:"completionTime,omitempty"`
	CompletionTokens *int     `json:"completionTokens,omitempty"`
	PromptTime       *float64 `json:"promptTime,omitempty"`
	PromptTokens     *int     `json:"promptTokens,omitempty"`
	QueueTime        *float64 `json:"queueTime,omitempty"`
	TotalTime        *float64 `json:"totalTime,omitempty"`
	TotalTokens      *int     `json:"totalTokens,omitempty"`
}
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
type GroqChatCompletionParams struct {
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

func (c *Client) GroqChatCompletion(ctx context.Context, req *GroqChatCompletionParams) (*GroqChatCompletionResponse, error) {
	if len(req.Messages) == 0 {
		return nil, ErrGroqChatCompletionNoMessage
	}
	if req.Model == "" {
		return nil, ErrModelNotFound
	}
	body, err := c.newJSONPostRequest(ctx, "/grok/chatcompletion", "", req)
	if err != nil {
		return nil, err
	}
	var groqChatCompletionResponse GroqChatCompletionResponse
	if err := json.Unmarshal(body, &groqChatCompletionResponse); err != nil {
		return nil, err
	}
	return &groqChatCompletionResponse, nil
}
