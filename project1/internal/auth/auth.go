package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts an API key from the request header
// Example:
// Authorization: ApiKey {apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("missing Authorization header")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid Authorization header format")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("invalid Authorization header type, expected 'ApiKey'")
	}
	return vals[1], nil
}
