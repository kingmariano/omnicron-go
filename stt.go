package omnicron

import (
	"context"
	"encoding/json"
	"os"
)

// model is ranked from lowest to highest based on their strength and abilites to perform Speech To Text (STT) generation. Check out the models on replicate.

type ReplicateLowSTTModel string

type ReplicateHighSTTModel string 

const (
	// Model on Replicate: https://replicate.com/openai/whisper
	WhisperModel ReplicateLowSTTModel = "openai/whisper"
    // Model on Replicate: https://replicate.com/turian/insanely-fast-whisper-with-video
	InsanelyFastWhisperWithVideoModel ReplicateHighSTTModel = "turian/insanely-fast-whisper-with-video"
)
type LowSTTParams struct {
	Audio               *os.File `json:"audio"`
	Transcription           *string         `json:"transcription,omitempty"`
	Temperature             *float64        `json:"temperature,omitempty"`
	Translate               *bool           `json:"translate,omitempty"`
	InitialPrompt           *string         `json:"initial_prompt,omitempty"`
	ConditionOnPreviousText *bool           `json:"condition_on_previous_text,omitempty"`
}
type HighSTTParams struct {
	AudioFile *os.File `json:"audio_file"`
	URL       *string         `json:"url,omitempty"`
	Task      *string         `json:"task,omitempty"`
	BatchSize *int            `json:"batch_size,omitempty"`
	Timestamp *string         `json:"timestamp,omitempty"`
}

type LowSTTModelAndParams struct{
	Model ReplicateLowSTTModel
	Parameters LowSTTParams
}

type HighSTTModelAndParams struct{
    Model ReplicateHighSTTModel
    Parameters HighSTTParams
}
// speech to text generation function
func (c *Client) LowSTTGeneration(ctx context.Context, req LowSTTModelAndParams) (*ReplicatePredictionResponse, error) {
	if req.Model == "" {
		return nil, ErrModelNotFound
	}
	body, err := c.newFormWithFilePostRequest(ctx, "/replicate/tts", string(req.Model), req.Parameters)
	if err != nil {
		return nil, err
	}
	var predictionResponse ReplicatePredictionResponse
	if err := json.Unmarshal(body, &predictionResponse); err != nil {
		return nil, err
	}
	return &predictionResponse, nil
}
func (c *Client) HighSTTGeneration(ctx context.Context, req HighSTTModelAndParams) (*ReplicatePredictionResponse, error) {
	if req.Model == "" {
		return nil, ErrModelNotFound
	}
	body, err := c.newFormWithFilePostRequest(ctx, "/replicate/tts", string(req.Model), req.Parameters)
	if err != nil {
		return nil, err
	}
	var predictionResponse ReplicatePredictionResponse
	if err := json.Unmarshal(body, &predictionResponse); err != nil {
		return nil, err
	}
	return &predictionResponse, nil
}