package middleware

import (
	"github.com/spf13/viper"
	"net/http"
)

func VerifyApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKeyHeader := r.Header.Get("X-API-KEY")
		apiKeyEnv := viper.GetString("API_KEY")
		if apiKeyHeader != apiKeyEnv {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
