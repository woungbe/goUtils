package chains

import (
	"log"
	"net/http"
)

// LoggingMiddleware는 각 요청에 대한 로깅을 수행하는 미들웨어입니다.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Request handled: %s %s", r.Method, r.URL.Path)
	})
}

// AuthMiddleware는 요청이 인증된 사용자로부터 왔는지 확인합니다.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
