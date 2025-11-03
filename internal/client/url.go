package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"go.uber.org/zap"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils/logging"
)

type Accept string

const (
	AcceptJSON = Accept("application/json")
	AcceptText = Accept("text/plain")
)

func BuildURL(baseURL string, pathParams []string, queryParams map[string][]string) (string, error) {
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

func BuildHttpRequest(method, baseURL string, pathParams []string, queryParams map[string][]string, body []byte, token string, accept Accept) (*http.Request, error) {
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

	if body != nil {
		req.Header.Set("Content-Type", string(AcceptJSON))
	}

	req.Header.Set("Authorization", "Bearer "+token)
	//  "application/json" or "text/plain"
	req.Header.Set("Accept", string(accept))

	return req, nil
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
		params[key] = slice
		return
	}

	// Handle pointer types
	switch v := value.(type) {
	case *string:
		if v != nil && *v != "" && *v != invalid {
			params.Set(key, *v)
		}
		return
	case *int:
		if v != nil && *v > 0 {
			params.Set(key, strconv.Itoa(*v))
		}
		return
	case *bool:
		if v != nil {
			params.Set(key, strconv.FormatBool(*v))
		}
		return
	default:
		// For all other types, compare with invalid value
		if value != invalid {
			// Special handling for basic types
			switch val := value.(type) {
			case string:
				if val != "" {
					params.Set(key, val)
				}
			case int:
				if val > 0 {
					params.Set(key, strconv.Itoa(val))
				}
			case bool:
				params.Set(key, strconv.FormatBool(val))
			default:
				if valStr := fmt.Sprintf("%v", val); valStr != "" {
					params.Set(key, valStr)
				}
			}
		}
	}
}

// SetRequestBodyParam sets a request body parameter based on its value.
// The parameter is set only if the value is not empty or zero.
// For string types, the parameter is set only if not empty.
// For slice types, the parameter is set only if the slice is not empty.
// For numeric types, the parameter is set only if it's not zero.
// For boolean types, the parameter is always set.
// For pointer types, the parameter is set only if the pointer is not nil.
// For MapOutput and []MapOutput types, the parameter is set only if not nil/empty.
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
	case types.MapOutput:
		// For MapOutput, set only if not nil
		if v != nil {
			params[key] = v
		}
	case []types.MapOutput:
		// For []MapOutput, set only if not empty
		if len(v) > 0 {
			params[key] = v
		}
	default:
		// For other types, only set if not nil
		if v != nil {
			params[key] = v
		}
	}
}

// ReadBody reads the response body and returns it as a string
func ReadBody(resp *http.Response) (string, error) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
