package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from the headers of an HTTP request
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authorization header not found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("authorization header has invalid format")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("authorization header has invalid format")
	}
	return vals[1], nil
}
