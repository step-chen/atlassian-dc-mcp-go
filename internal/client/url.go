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

// convertAnyToStringSlice converts a slice of any to a slice of strings.
// It preserves the order of elements and converts each element to its string representation.
func convertAnyToStringSlice(slice []any) []string {
	if slice == nil {
		return nil
	}

	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}

type Accept string

const (
	AcceptJSON = Accept("application/json")
	AcceptText = Accept("text/plain")
)

func buildURL(baseURL string, pathSegments []any, queryParams map[string][]string) (string, error) {
	stringPathParams := convertAnyToStringSlice(pathSegments)
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse base URL: %w", err)
	}

	u = u.JoinPath(stringPathParams...)

	q := u.Query()
	for key, values := range queryParams {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func buildHttpRequest(method, baseURL string, pathSegments []any, queryParams map[string][]string, body []byte, token string, accept Accept) (*http.Request, error) {
	url, err := buildURL(baseURL, pathSegments, queryParams)
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

func SetRequiredPathParam(queryParams url.Values, path string) {
	if path == "" {
		queryParams.Add("path", "/")
	} else {
		queryParams.Add("path", path)
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
func SetQueryParam(queryParams url.Values, key string, value any, invalid any) bool {
	switch v := value.(type) {
	case []string:
		if len(v) == 0 {
			return false
		}
		queryParams[key] = v
		return true
	case *string:
		if v != nil && *v != "" && *v != invalid {
			queryParams.Set(key, *v)
			return true
		}
	case *int:
		if v != nil && *v > 0 {
			queryParams.Set(key, strconv.Itoa(*v))
			return true
		}
	case *bool:
		if v != nil {
			queryParams.Set(key, strconv.FormatBool(*v))
			return true
		}
	case string:
		if v != "" && v != invalid {
			queryParams.Set(key, v)
			return true
		}
	case int:
		if v > 0 && v != invalid {
			queryParams.Set(key, strconv.Itoa(v))
			return true
		}
	case bool:
		queryParams.Set(key, strconv.FormatBool(v))
		return true
	default:
		if value != invalid {
			if valStr := fmt.Sprintf("%v", value); valStr != "" && valStr != fmt.Sprintf("%v", invalid) {
				queryParams.Set(key, valStr)
				return true
			}
		}
	}
	return false
}

// SetRequestBodyParam sets a request body parameter based on its value.
// The parameter is set only if the value is not empty or zero.
// For string types, the parameter is set only if not empty.
// For slice types, the parameter is set only if the slice is not empty.
// For numeric types, the parameter is set only if it's not zero.
// For boolean types, the parameter is always set.
// For pointer types, the parameter is set only if the pointer is not nil.
// For MapOutput and []MapOutput types, the parameter is set only if not nil/empty.
func SetRequestBodyParam(bodyParams map[string]interface{}, key string, value interface{}) {
	switch v := value.(type) {
	case string:
		if v != "" {
			bodyParams[key] = v
		}
	case []string:
		if len(v) > 0 {
			bodyParams[key] = v
		}
	case int, int8, int16, int32, int64:
		if v != 0 {
			bodyParams[key] = v
		}
	case float32, float64:
		if v != 0.0 {
			bodyParams[key] = v
		}
	case bool:
		bodyParams[key] = v
	case *string:
		if v != nil && *v != "" {
			bodyParams[key] = *v
		}
	case *int:
		if v != nil && *v != 0 {
			bodyParams[key] = *v
		}
	case *bool:
		if v != nil {
			bodyParams[key] = *v
		}
	case types.MapOutput:
		if v != nil {
			bodyParams[key] = v
		}
	case []types.MapOutput:
		if len(v) > 0 {
			bodyParams[key] = v
		}
	default:
		if v != nil {
			bodyParams[key] = v
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
