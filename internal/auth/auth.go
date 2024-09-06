package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts API Key from request headers
// Example:
// Authorization: ApiKey {insert API Key here}
func GetAPIKey(headers http.Header) (string, error) {
	// Get the header value
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth info found")
	}

	// Check if both parts of header exist
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	// Check if first part of header is written correctly
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}
