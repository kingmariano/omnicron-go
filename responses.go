package omnicron

type Status string

type Source string

type WebhookEventType string

const (
	WebhookEventStart     WebhookEventType = "start"
	WebhookEventOutput    WebhookEventType = "output"
	WebhookEventLogs      WebhookEventType = "logs"
	WebhookEventCompleted WebhookEventType = "completed"
)

const (
	SourceWeb Source = "web"
	SourceAPI Source = "api"
)

const (
	Starting   Status = "starting"
	Processing Status = "processing"
	Succeeded  Status = "succeeded"
	Failed     Status = "failed"
	Canceled   Status = "canceled"
)

type PredictionMetrics struct {
	PredictTime      *float64 `json:"predict_time,omitempty"`
	TotalTime        *float64 `json:"total_time,omitempty"`
	InputTokenCount  *int     `json:"input_token_count,omitempty"`
	OutputTokenCount *int     `json:"output_token_count,omitempty"`
	TimeToFirstToken *float64 `json:"time_to_first_token,omitempty"`
	TokensPerSecond  *float64 `json:"tokens_per_second,omitempty"`
}
type PredictionInput map[string]interface{}
type PredictionOutput interface{}
type ReplicatePredictionResponse struct {
	ID                  string             `json:"id"`
	Status              Status             `json:"status"`
	Model               string             `json:"model"`
	Version             string             `json:"version"`
	Input               PredictionInput    `json:"input"`
	Output              PredictionOutput   `json:"output,omitempty"`
	Source              Source             `json:"source"`
	Error               interface{}        `json:"error,omitempty"`
	Logs                *string            `json:"logs,omitempty"`
	Metrics             *PredictionMetrics `json:"metrics,omitempty"`
	Webhook             *string            `json:"webhook,omitempty"`
	WebhookEventsFilter []WebhookEventType `json:"webhook_events_filter,omitempty"`
	URLs                map[string]string  `json:"urls,omitempty"`
	CreatedAt           string             `json:"created_at"`
	StartedAt           *string            `json:"started_at,omitempty"`
	CompletedAt         *string            `json:"completed_at,omitempty"`

	// rawJSON json.RawMessage `json:"-"`
}

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

type Responseparams struct {
	Response string `json:"response"`
}