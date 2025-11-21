// Package utils provides utility functions for HTTP clients and other common operations.
package client

import (
	"fmt"
	"regexp"
	"strings"

	"atlassian-dc-mcp-go/internal/config"
)

var (
	fuzzyKeys   []string
	removePaths []string
)

// InitPruneConfig initializes the prune configuration from the config
func InitPruneConfig(cfg config.PruneConfig) {
	fuzzyKeys = cfg.FuzzyKeys
	removePaths = cfg.RemovePaths
}

func Prune(m any) {
	for {
		switch v := m.(type) {
		case *map[string]any:
			if v == nil {
				return
			}
			m = *v
		case *[]any:
			if v == nil {
				return
			}
			m = *v
		default:
			goto done
		}
	}
done:
	switch m := m.(type) {
	case map[string]any:
		prune(m, "")
	case []any:
		for i, v := range m {
			if v, ok := v.(map[string]any); ok {
				prune(v, fmt.Sprintf("[%d]", i))
			}
		}
	}
}

func prune(m map[string]any, prefix string) {
	for k, v := range m {
		currentPath := k
		if prefix != "" {
			currentPath = prefix + "." + k
		}

		if shouldRemove(currentPath) {
			delete(m, k)
			continue
		}

		if isZeroValue(v) {
			delete(m, k)
			continue
		}

		switch vv := v.(type) {
		case map[string]any:
			prune(vv, currentPath)
		case []any:
			for i, item := range vv {
				if itemMap, ok := item.(map[string]any); ok {
					prune(itemMap, currentPath+fmt.Sprintf("[%d]", i))
				}
			}
		}
	}
}

func isZeroValue(v any) bool {
	switch vv := v.(type) {
	case nil:
		return true
	case string:
		return vv == ""
	case map[string]any:
		return len(vv) == 0
	case []any:
		return len(vv) == 0
	default:
		return false
	}
}

func fuzzyMatch(key string) bool {
	for _, fk := range fuzzyKeys {
		if strings.HasPrefix(key, fk) {
			return true
		}
	}
	return false
}

func shouldRemove(path string) bool {
	arrayIndexRegex := regexp.MustCompile(`\[\d+\]`)
	cleanPath := arrayIndexRegex.ReplaceAllString(path, "")

	for _, rp := range removePaths {
		if strings.HasSuffix(cleanPath, "."+rp) || cleanPath == rp {
			return true
		}
	}

	keys := strings.Split(cleanPath, ".")
	if len(keys) > 0 {
		lastKey := keys[len(keys)-1]
		if fuzzyMatch(lastKey) {
			return true
		}
	}
	return false
}
