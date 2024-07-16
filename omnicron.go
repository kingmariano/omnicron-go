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
	"reflect"
	"strconv"
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
func (c *Client) newJSONPostRequest(ctx context.Context, path, model string, payload interface{}) ([]byte, error) {
	fullURLPath := c.baseurl + "api/v1" + path
	if model != "" {
		fullURLPath = c.withModelQueryParameters(fullURLPath, model)
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
func (c *Client) newFormWithFilePostRequest(ctx context.Context, path, model string, payload interface{}) ([]byte, error) {
	fullURLPath := c.baseurl + path
	if model != "" {
		fullURLPath = c.withModelQueryParameters(fullURLPath, model)
	}
	if c.debug {
		log.Printf("full url path: %s", fullURLPath)
	}

	// Create a buffer to hold the form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	v := reflect.ValueOf(payload)
	typeOfParams := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := typeOfParams.Field(i).Tag.Get("form")

		if fieldName == "" {
			continue
		}

		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		switch field.Interface().(type) {
		case *os.File:
			if fieldName == "image" || fieldName == "mask" {
				addFileField(writer, fieldName, field.Interface().(*os.File))
			}
		case *int:
			addField(writer, fieldName, strconv.Itoa(*field.Interface().(*int)))
		case *float64:
			addField(writer, fieldName, fmt.Sprintf("%f", *field.Interface().(*float64)))
		case *string:
			addField(writer, fieldName, *field.Interface().(*string))
		case *bool:
			addField(writer, fieldName, strconv.FormatBool(*field.Interface().(*bool)))
		default:
			addField(writer, fieldName, fmt.Sprintf("%v", field.Interface()))
		}
	}
	writer.Close()
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
func addField(writer *multipart.Writer, key string, value string) error {
	err := writer.WriteField(key, value)
	if err != nil {
		return err
	}
	return nil
}

func addFileField(writer *multipart.Writer, fieldname string, file *os.File) error {
	defer file.Close()
	fw, err := writer.CreateFormFile(fieldname, file.Name())
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, file)
	return err
}

func (c *Client) withModelQueryParameters(fullURLPath, model string) string {
	params := url.Values{}
	params.Add("model", model)
	url := fullURLPath + "?" + params.Encode()
	return url
}
