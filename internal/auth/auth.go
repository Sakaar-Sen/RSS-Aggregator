package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("api key not found")
	}

	valfields := strings.Fields(val)
	if len(valfields) != 2 {
		return "", errors.New("invalid API key")
	}

	if valfields[0] != "Bearer" {
		return "", errors.New("invalid API key")
	}

	return valfields[1], nil
}
