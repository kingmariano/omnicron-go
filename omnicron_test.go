package omnicron

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

// MockHTTPClient is a mock implementation of the HTTPClient interface
type mockHTTPClient struct {
	Response []byte
	Err      error
}

func (c *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if c.Err != nil {
		return nil, c.Err
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(c.Response)),
	}, nil
}

// NewMockClient creates a new Client with a mock HTTP client
func NewMockClient(response []byte, err error) *Client {
	mockClient := &mockHTTPClient{Response: response, Err: err}
	return &Client{
		apikey:  "your-api-key",
		baseurl: "https://custom-url.com",
		httpClient:  mockClient,
	}
}

func TestGrokChatCompletion(t *testing.T) {
	type args struct {
		model    string
		messages []Message
	}
	testCases := []struct {
		name          string
		args          args
		mockResponse  []byte
		mockError     error
		expectedError bool
	}{
		{
			name: "successful request",
			args: args{
				model: "llama3-8b-8192",
				messages: []Message{
					{
						Role:    "user",
						Content: "This is a test message",
					},
				},
			},
			mockResponse:  []byte(`{"choices":[{"message":{"role":"assistant","content":"Test response"}}]}`),
			expectedError: false,
		},
		{
			name: "error when no model is provided",
			args: args{
				model: "",
				messages: []Message{
					{
						Role:    "user",
						Content: "This is a test message",
					},
				},
			},
			mockResponse:  []byte(`{"error":"No model provided"}`),
			expectedError: true,
		},
		{
			name: "no message provided",
			args: args{
				model:    "llama3-8b-8192",
				messages: []Message{},
			},
			mockResponse:  []byte(`{"error":"No message provided"}`),
			expectedError: true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			client := NewMockClient(tt.mockResponse, tt.mockError)
			ctx := context.Background()
           req := &GroqChatCompletionParams{
             Messages: tt.args.messages,
			 Model: tt.args.model,
		   }
			_, err := client.GroqChatCompletion(ctx, req)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
