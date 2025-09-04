package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func BuildURL(baseURL string, pathParams []string, queryParams url.Values) (string, error) {
	fullPath, err := url.JoinPath(baseURL, pathParams...)
	if err != nil {
		return "", fmt.Errorf("failed to join path segments: %w", err)
	}

	if len(queryParams) > 0 {
		return fullPath + "?" + queryParams.Encode(), nil
	}

	return fullPath, nil
}

func BuildHttpRequest(method, baseURL string, pathParams []string, queryParams url.Values, body []byte, token string) (*http.Request, error) {
	url, err := BuildURL(baseURL, pathParams, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL: %w", err)
	}

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

func SetQueryParam(params url.Values, key string, value any, invalid any) {
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
	case []string:
		if len(v) > 0 {
			for _, s := range v {
				params.Add(key, s)
			}
		}
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