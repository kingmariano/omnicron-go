/*
This music generation function uses different and high efficient music generation AI models on Replicate. Please note that cost varies as you some of this model.
*/
package omnicron

import (
	"context"
	"encoding/json"
	"os"
)

// These models is described low cause it doesnt support an input audio file for the music generation
type ReplicateLowMusicGenerationModel string

// These models is described high cause it supports an input audio file for the music generation
type ReplicateHighMusicGenerationModel string

const (

	// Model on Replicate: https://replicate.com/riffusion/riffusion
	RiffusionModel ReplicateLowMusicGenerationModel = "riffusion/riffusion"

	// Model on Replicate: https://replicate.com/meta/musicgen
	MetaMusicGenModel ReplicateHighMusicGenerationModel = "meta/musicgen"
)

// LowMusicGenerationParams groups the low music generation model with its parameters.
type LowMusicGenerationParams struct {
	PromptA           string   `json:"prompt_a"`
	Denoising         *float64 `json:"denoising,omitempty"`
	PromptB           *string  `json:"prompt_b,omitempty"`
	Alpha             *float64 `json:"alpha,omitempty"`
	NumInferenceSteps *int     `json:"num_inference_steps,omitempty"`
	SeedImageID       *string  `json:"seed_image_id,omitempty"`
}

// HighMusicGenerationParams groups the high music generation model with its parameters
type HighMusicGenerationParams struct {
	Prompt                 string   `form:"prompt"`
	ModelVersion           *string  `form:"model_version,omitempty"`
	InputAudioFile         *os.File `form:"input_audio,omitempty"`
	Duration               *int     `form:"duration,omitempty"`
	Continuation           *bool    `form:"continuation,omitempty"`
	ContinuationStart      *int     `form:"continuation_start,omitempty"`
	ContinuationEnd        *int     `form:"continuation_end,omitempty"`
	MultiBandDiffusion     *bool    `form:"multi_band_diffusion,omitempty"`
	NormalizationStrategy  *string  `form:"normalization_strategy,omitempty"`
	TopK                   *int     `form:"top_k,omitempty"`
	TopP                   *float64 `form:"top_p,omitempty"`
	Temperature            *float64 `form:"temperature,omitempty"`
	ClassifierFreeGuidance *int     `form:"classifier_free_guidance,omitempty"`
	OutputFormat           *string  `form:"output_format,omitempty"`
}

// LowMusicGenerationModelAndParams groups the low image generation model with its parameters.
type LowMusicGenerationModelAndParams struct {
	Model      ReplicateLowMusicGenerationModel
	Parameters *LowMusicGenerationParams
}

// HighMusicGenerationModelAndParams groups the high image generation model with its parameters.
type HighMusicGenerationModelAndParams struct {
	Model      ReplicateHighMusicGenerationModel
	Parameters HighMusicGenerationParams
}

func (c *Client) LowMusicGeneration(ctx context.Context, req LowMusicGenerationModelAndParams) (*ReplicatePredictionResponse, error) {
	if req.Model == "" {
		return nil, ErrModelNotFound
	}
	body, err := c.newJSONPostRequest(ctx, "/replicate/musicgeneration", string(req.Model), req.Parameters)
	if err != nil {
		return nil, err
	}
	var predictionResponse ReplicatePredictionResponse
	if err := json.Unmarshal(body, &predictionResponse); err != nil {
		return nil, err
	}
	return &predictionResponse, nil
}
func (c *Client) HighMusicGeneration(ctx context.Context, req HighMusicGenerationModelAndParams) (*ReplicatePredictionResponse, error) {
	if req.Model == "" {
		return nil, ErrModelNotFound
	}
	if req.Parameters.Prompt == "" {
		return nil, ErrPromptMissing
	}

	body, err := c.newFormWithFilePostRequest(ctx, "/replicate/musicgeneration", string(req.Model), req.Parameters)
	if err != nil {
		return nil, err
	}
	var predictionResponse ReplicatePredictionResponse
	if err := json.Unmarshal(body, &predictionResponse); err != nil {
		return nil, err
	}
	return &predictionResponse, nil
}
