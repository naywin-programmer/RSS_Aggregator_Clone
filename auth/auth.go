package auth

import (
	"errors"
	"net/http"
	"strings"
)

const UserApiKeyPrefixInHeader = "ApiKey"

// example header format
// Authorization: {apiKeyPrefixInHeader} {user_api_key_in_here}
func GetApiKey(header http.Header, apiKeyPrefixInHeader string) (string, error) {
	requestedApiKeyHeader := header.Get("Authorization")
	if requestedApiKeyHeader == "" {
		return "", errors.New("no Authorization Info Found")
	}

	apiKeyHeaderArr := strings.Split(requestedApiKeyHeader, " ")
	if len(apiKeyHeaderArr) != 2 {
		return "", errors.New("invalid Authorization Format")
	}

	if apiKeyHeaderArr[0] != apiKeyPrefixInHeader {
		return "", errors.New("invalid Authorization Prefix Format")
	}

	return apiKeyHeaderArr[1], nil
}
