package main

import (
	"fmt"
	"log"
	"net/http"
)

// 인터페이스 설정
type Middleware interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

// 인증 미들웨어 설정
type AuthMiddleware struct{}

func (a AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// 인증 로직
	if !isAuthenticated(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	next(w, r)
}

// 실제 사용
func WithMiddleware(h http.HandlerFunc, m Middleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.ServeHTTP(w, r, h)
	}
}

type CheckAdminUserMiddlewere struct{}

func (a CheckAdminUserMiddlewere) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// 인증 로직
	if !isAdmin(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 하고 싶은거 다 하면 되겠지 ...
	// 유저한테 뭘 넣은 다던가.
	// IP 차단을 한다던가.
	// 특정 유저의 수수료를 올린다던가, 내린다던가.
	// 어떤 선검사 같은 것도 괜찮겠지.. 전체적으로 싹다.
	//

	next(w, r)
}

func isAdmin(r *http.Request) bool {
	token := r.Header.Get("I am Admin")
	if token == "admin" {
		return true
	}
	return false
}

func isAuthenticated(r *http.Request) bool {
	// 예를 들어, 헤더에서 토큰을 추출
	token := r.Header.Get("Authorization")
	// 토큰 검증 로직
	return token == "some_valid_token"
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	// 실제 핸들러 로직, 예를 들어:
	fmt.Fprintf(w, "Hello, authenticated user!")
}

func main() {

	// 아 실제 - Middleware 이건 그냥 받아야되는 변수 설정이구나 ..
	http.HandleFunc("/", WithMiddleware(myHandler, AuthMiddleware{}))
	http.HandleFunc("/isAdmin", WithMiddleware(myHandler, CheckAdminUserMiddlewere{}))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
