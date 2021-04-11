package handlers

import (
	"fmt"
	"introGO/examples/basics/test/cache"
	"log"
	"net/http"
	"strings"
)

func TokenMiddleware(tokenCache cache.TokenCacheService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("Authorization")

		if len(header) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			log.Println("Missing Authorization Header")
			return
		}

		tokenString := strings.Replace(header, "Bearer ", "", 1)

		fmt.Printf("Token is %s \n", tokenString)

		expired, err := tokenCache.IsExpired(tokenString)

		if err != nil || expired {
			http.Error(w, "Not authorized", 401)
			return
		}

		next.ServeHTTP(w, r)
	})
}
