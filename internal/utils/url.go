package utils

import (
	"atlassian-dc-mcp-go/internal/utils/logging"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"go.uber.org/zap"
)

func BuildURL(baseURL string, pathParams []string, queryParams url.Values) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse base URL: %w", err)
	}

	u = u.JoinPath(pathParams...)

	q := u.Query()
	for key, values := range queryParams {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func BuildHttpRequest(method, baseURL string, pathParams []string, queryParams url.Values, body []byte, token string) (*http.Request, error) {
	url, err := BuildURL(baseURL, pathParams, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL: %w", err)
	}

	logging.GetLogger().Info("Building request", zap.String("method", method), zap.String("url", url))
	var req *http.Request
	if body != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	if body != nil && (method == http.MethodPost || method == http.MethodPut) {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func HandleHTTPError(resp *http.Response, service string) error {
	strErr := ""
	switch resp.StatusCode {
	case http.StatusBadRequest:
		strErr = "bad request"
	case http.StatusUnauthorized:
		strErr = "unauthorized"
	case http.StatusForbidden:
		strErr = "forbidden"
	case http.StatusNotFound:
		strErr = "not found"
	default:
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return nil
		} else {
			strErr = "unknown error"
		}
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	return fmt.Errorf("[%s] %s : %d - %s", service, strErr, resp.StatusCode, bodyString)
}

func SetRequiredPathQueryParam(params url.Values, path string) {
	if path == "" {
		params.Add("path", "/")
	} else {
		params.Add("path", path)
	}
}

// SetQueryParam sets a query parameter based on its value and an invalid value to compare against.
// If the value matches the invalid value, the parameter is not set.
// For slices, the parameter is set only if the slice is not empty.
// For other types, specific rules apply:
// - int: set only if greater than 0
// - string: set only if not empty
// - bool: always set
// - pointers: set only if not nil and the dereferenced value is valid
func SetQueryParam(params url.Values, key string, value any, invalid any) {
	// Handle slice comparisons specially since slices are not comparable
	if slice, ok := value.([]string); ok {
		// Check if slice is empty
		if len(slice) == 0 {
			// Also check if the invalid value is an empty slice
			if invalidSlice, ok := invalid.([]string); ok && len(invalidSlice) == 0 {
				return
			} else if invalid == nil {
				return
			}
		}

		// If we get here, the slice is not empty, so process it
		for _, s := range slice {
			params.Add(key, s)
		}
		return
	}

	// For non-slice types, use the original comparison
	if value == invalid {
		return
	}

	switch v := value.(type) {
	case int:
		if v > 0 {
			params.Set(key, strconv.Itoa(v))
		}
	case string:
		if v != "" {
			params.Set(key, v)
		}
	case bool:
		params.Set(key, strconv.FormatBool(v))
	case *string:
		if v != nil && *v != "" {
			params.Set(key, *v)
		}
	case *int:
		if v != nil && *v > 0 {
			params.Set(key, strconv.Itoa(*v))
		}
	case *bool:
		if v != nil {
			params.Set(key, strconv.FormatBool(*v))
		}
	default:
		if valStr := fmt.Sprintf("%v", v); valStr != "" {
			params.Set(key, valStr)
		}
	}
}

// SetRequestBodyParam sets a request body parameter based on its value.
// The parameter is set only if the value is not empty or zero.
// For string types, the parameter is set only if not empty.
// For numeric types, the parameter is set only if greater than zero.
// For boolean types, the parameter is always set.
// For pointer types, the parameter is set only if the pointer is not nil.
func SetRequestBodyParam(params map[string]interface{}, key string, value interface{}) {
	switch v := value.(type) {
	case string:
		if v != "" {
			params[key] = v
		}
	case []string:
		if len(v) > 0 {
			params[key] = v
		}
	case int, int8, int16, int32, int64:
		if v != 0 {
			params[key] = v
		}
	case float32, float64:
		if v != 0.0 {
			params[key] = v
		}
	case bool:
		params[key] = v
	case *string:
		if v != nil && *v != "" {
			params[key] = *v
		}
	case *int:
		if v != nil && *v != 0 {
			params[key] = *v
		}
	case *bool:
		if v != nil {
			params[key] = *v
		}
	default:
		// For other types, only set if not nil
		if v != nil {
			params[key] = v
		}
	}
}

func ExecuteHTTPRequest(client *http.Client, req *http.Request, service string, result interface{}) error {
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[%s] request failed: %w", service, err)
	}
	defer resp.Body.Close()

	if err := HandleHTTPError(resp, service); err != nil {
		return err
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("[%s] failed to decode response: %w", service, err)
		}
	}

	return nil
}

func HandleHTTPResponse(resp *http.Response, service string, result interface{}) error {
	defer resp.Body.Close()

	if err := HandleHTTPError(resp, service); err != nil {
		return err
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("[%s] failed to decode response: %w", service, err)
		}
	}

	return nil
}

// ReadBody reads the response body and returns it as a string
func ReadBody(resp *http.Response) (string, error) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
