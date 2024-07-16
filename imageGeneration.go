package omnicron

import (
	"context"
	"encoding/json"
	"os"
)

// image generations supports multiple image generation AI model on replicate.
// defines the image generation models on replicate
type ReplicateLowImageGenerationModel string  //doesn't support image to image processing
type ReplicateHighImageGenerationModel string //supports image to image processing
const (
	//model on replicate: https://replicate.com/bytedance/sdxl-lightning-4step
	SDXLLightning4stepModel ReplicateLowImageGenerationModel = "bytedance/sdxl-lightning-4step"
	//model on replicate: https://replicate.com/lucataco/realvisxl-v2.0
	RealvisxlV20Model ReplicateHighImageGenerationModel = "lucataco/realvisxl-v2.0"
	//model on replicate: https://replicate.com/playgroundai/playground-v2.5-1024px-aesthetic
	PlaygroundV251024pxAestheticModel ReplicateHighImageGenerationModel = "playgroundai/playground-v2.5-1024px-aesthetic"
	//model on replicate: https://replicate.com/lucataco/dreamshaper-xl-turbo
	DreamshaperXlTurboModel ReplicateLowImageGenerationModel = "lucataco/dreamshaper-xl-turbo"
	//model on replicate: https://replicate.com/lorenzomarines/astra
	AstraModel ReplicateHighImageGenerationModel = "lorenzomarines/astra"
)

// low image generation model because it doesn't support image to image processing
type LowImageGenerationParams struct {
	Prompt            string   `json:"prompt"`
	Width             *int     `json:"width,omitempty"`
	Height            *int     `json:"height,omitempty"`
	Scheduler         *string  `json:"scheduler,omitempty"`
	NumOutputs        *int     `json:"num_outputs,omitempty"`
	GuidanceScale     *float64 `json:"guidance_scale,omitempty"`
	NegativePrompt    *string  `json:"negative_prompt,omitempty"`
	NumInferenceSteps *int     `json:"num_inference_steps,omitempty"`
}
type LowImageGenerationModelAndParams struct {
	Model      ReplicateLowImageGenerationModel
	Parameters *LowImageGenerationParams
}

// high image generation model because it supports image to image processing
type HighImageGenerationParams struct {
	Prompt            string   `json:"prompt"`
	Width             *int     `json:"width,omitempty"`
	Height            *int     `json:"height,omitempty"`
	Scheduler         *string  `json:"scheduler,omitempty"`
	NumOutputs        *int     `json:"num_outputs,omitempty"`
	GuidanceScale     *float64 `json:"guidance_scale,omitempty"`
	NegativePrompt    *string  `json:"negative_prompt,omitempty"`
	NumInferenceSteps *int     `json:"num_inference_steps,omitempty"`
	LoraScale         *float64 `json:"lora_scale,omitempty"`
	ImageFile         *os.File `json:"image_file,omitempty"`
	MaskFile          *os.File `json:"mask_file,omitempty"`
	PromptStrength    *float64 `json:"prompt_strength,omitempty"`
	ApplyWatermark    *bool    `json:"apply_watermark,omitempty"`
	Seed              *int     `json:"seed,omitempty"`
}

type HighImageGenerationModelAndParams struct {
	Model      ReplicateHighImageGenerationModel
	Parameters *HighImageGenerationParams
}

func (c *Client) LowImageGeneration(ctx context.Context, req LowImageGenerationModelAndParams) (*ReplicatePredictionResponse, error) {
	if req.Model == "" {
		return nil, ErrModelNotFound
	}
	body, err := c.newJSONPostRequest(ctx, "/replicate/imagegeneration", string(req.Model), req.Parameters)
	if err != nil {
		return nil, err
	}
	var predictionResponse ReplicatePredictionResponse
	if err := json.Unmarshal(body, &predictionResponse); err != nil {
		return nil, err
	}
	return &predictionResponse, nil
}

func (c *Client) HighImageGeneration(ctx context.Context, req HighImageGenerationModelAndParams) (*ReplicatePredictionResponse, error) {
	if req.Model == "" {
		return nil, ErrModelNotFound
	}
	if req.Parameters.Prompt == "" {
		return nil, ErrPromptMissing
	}

	body, err := c.newFormWithFilePostRequest(ctx, "/replicate/imagegeneration", string(req.Model), req.Parameters)
	if err != nil {
		return nil, err
	}
	var predictionResponse ReplicatePredictionResponse
	if err := json.Unmarshal(body, &predictionResponse); err != nil {
		return nil, err
	}
	return &predictionResponse, nil
}
