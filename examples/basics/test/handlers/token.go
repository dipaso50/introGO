package handlers

import (
	"encoding/json"
	"introGO/examples/basics/test/cache"
	"net/http"

	"github.com/google/uuid"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func CreateToken(tokenCache cache.TokenCacheService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()

		//add token to cache
		tokenCache.AddToken(cache.Token{Value: id.String()})

		tokenResp := TokenResponse{id.String()}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenResp)
	})
}
