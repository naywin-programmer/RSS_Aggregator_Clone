package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	cf "github.com/naywin-programmer/RSS_Aggregator/configs"
	rs "github.com/naywin-programmer/RSS_Aggregator/resources"
	ut "github.com/naywin-programmer/RSS_Aggregator/utils"
)

// type basicApiAuthContextKey string
// const basicApiAuthContextKeyString basicApiAuthContextKey = "basic_api_auth_user"

type basicApiAuthMiddleware struct {
	AuthUser rs.UserResource
	Error    error
}

var BasicApiAuthMiddleware = basicApiAuthMiddleware{}

func (t *basicApiAuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKeyPrefixInHeader := "API_KEY"

		apiKey, err := t.GetApiKeyFromHeader(r.Header, apiKeyPrefixInHeader)

		if err != nil {
			t.AuthUser = rs.UserResource{}
			t.Error = err
			ut.RespondErrorWithJSON(w, http.StatusForbidden, fmt.Sprintf("authentication error: %v", err))
			return
		}

		user, err := cf.DBConnection.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			t.AuthUser = rs.UserResource{}
			t.Error = err
			ut.RespondErrorWithJSON(w, http.StatusBadRequest, fmt.Sprintf("user is not found: %v", err))
			return
		}

		t.AuthUser = rs.ToUserResource(user)
		t.Error = nil

		// ctx := context.WithValue(r.Context(), basicApiAuthContextKeyString, t)
		// next.ServeHTTP(w, r.WithContext(ctx))

		next.ServeHTTP(w, r)
	})
}

func (t *basicApiAuthMiddleware) GetApiKeyFromHeader(header http.Header, apiKeyPrefixInHeader string) (string, error) {
	// Authorization: {apiKeyPrefixInHeader} {user_api_key_in_here}
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
