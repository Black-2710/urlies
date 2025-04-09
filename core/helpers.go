// core/helpers.go
package core

import (
	"net/url"
	"strings"
)

// IsValidURL checks if a string is a valid URL
func IsValidURL(u string) bool {
	parsed, err := url.ParseRequestURI(u)
	return err == nil && parsed.Scheme != "" && parsed.Host != ""
}

// CleanURL trims tracking parameters, etc.
func CleanURL(raw string) string {
	if i := strings.Index(raw, "&"); i != -1 {
		return raw[:i]
	}
	return raw
}

// ExtractDomain extracts hostname from URL
func ExtractDomain(raw string) string {
	parsed, err := url.Parse(raw)
	if err != nil {
		return "unknown"
	}
	return parsed.Hostname()
}
