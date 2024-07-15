package omnicron

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	baseurl    string
	apikey     string
	debug      bool
	httpClient HTTPClient
}

type ClientOption func(*Client)

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseurl = baseURL
	}
}
func WithDebug(debug bool) ClientOption {
	return func(c *Client) {
		c.debug = debug
	}
}

func NewClient(apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		apikey:  apiKey,
		debug:   false,
		baseurl: "https://omnicron-latest.onrender.com/", //default base url for omnicron runs on 0.1 CPU 512 MB Ram
	}
	for _, opt := range opts {
		opt(c)
	}
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	return c
}
func (c *Client) newJSONPostRequest(ctx context.Context, path, query string, payload interface{}) ([]byte, error) {
	fullURLPath := c.baseurl + "api/v1" + path
	if query != "" {
		fullURLPath = c.withQueryParameters(fullURLPath, query)
	}
	//debug set
	if c.debug {
		log.Printf("full url path: %s", fullURLPath)
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	if c.debug {
		log.Println(string(body))
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURLPath, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if c.apikey != "" {
		httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apikey))
	}
	res, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if c.debug {
		log.Println(string(resBody))
	}
	if res.StatusCode != http.StatusOK {
		errResp := ErrorResponse{}
		if err := json.Unmarshal(resBody, &errResp); err != nil {
			return nil, fmt.Errorf("error unmarshalling: %s", resBody)
		}
		return nil, fmt.Errorf("API request failed: %s", errResp.Error)
	}
	return resBody, nil
}
func (c *Client) newFormWithFilePostRequest(ctx context.Context, path, query string, formFields map[string]string, fileFields map[string]*os.File) ([]byte, error) {
	fullURLPath := c.baseurl + path
	if query != "" {
		fullURLPath = c.withQueryParameters(fullURLPath, query)
	}
	if c.debug {
		log.Printf("full url path: %s", fullURLPath)
	}

	// Create a buffer to hold the form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add form fields to the form data
	for key, value := range formFields {
		err := writer.WriteField(key, value)
		if err != nil {
			return nil, err
		}
	}

	// Add file fields to the form data
	for fieldname, file := range fileFields {
		fileWriter, err := writer.CreateFormFile(fieldname, file.Name())
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			return nil, err
		}
	}

	// Close the writer to finalize the form data
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	// Create a new POST request with the form data
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURLPath, body)
	if err != nil {
		return nil, err
	}

	// Set the content type to multipart/form-data
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())

	// Set the Authorization header if an API key is provided
	if c.apikey != "" {
		httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apikey))
	}

	// Send the request
	res, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Read the response body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Log the response body if debug mode is enabled
	if c.debug {
		log.Println(string(resBody))
	}

	// Check for non-200 status codes and handle the error response
	if res.StatusCode != http.StatusOK {
		errResp := ErrorResponse{}
		if err := json.Unmarshal(resBody, &errResp); err != nil {
			return nil, fmt.Errorf("error unmarshalling: %s", resBody)
		}
		return nil, fmt.Errorf("API request failed: %s", errResp.Error)
	}

	return resBody, nil
}

// Ptr returns a pointer to the given value.
func Ptr[T any](v T) *T {
	return &v
}

func (c *Client) withQueryParameters(fullURLPath, query string) string {
	params := url.Values{}
	params.Add("model", query)
	url := fullURLPath + "?" + params.Encode()
	return url
}
