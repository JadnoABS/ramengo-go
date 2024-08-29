package middleware

import (
	"os"
	"net/http"
)

type ApiKeyMiddleware struct {
	ApiKey string
}

func NewApiKeyMiddleware() *ApiKeyMiddleware {
	var key = os.Getenv("RV_API_KEY")
	return &ApiKeyMiddleware{ ApiKey: key }
}

func (middleware *ApiKeyMiddleware) Apply(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return middleware.applyMiddleware(http.HandlerFunc(handler))
}

func (middleware *ApiKeyMiddleware) applyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		requested := r.Header.Get("x-api-key")
		if requested == "" || requested != middleware.ApiKey {
			http.Error(w, "x-api-key header missing", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}