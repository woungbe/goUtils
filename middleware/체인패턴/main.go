package main

import (
	"fmt"
	"log"
	"net/http"

	chains "study.middleware.chain/chains"
)

// 반복할 func 형태
// http.Handler 도 일종의 interface 형태임 - 그러니까 체인에서는 interface 형이 쓰는 형태.
type Middleware func(http.Handler) http.Handler

// 반복처리할 func
func ChainList(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// 이미 만들어져 있는 func 에 대해서도 체인을 걸어서 사용이 가능하다.
// 물론 reqeust , response 가 좀 달라지겠지만. 그래도 상관없음.....
func main() {
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// 미들웨어 체인을 구성
	chainedHandler := ChainList(finalHandler, chains.LoggingMiddleware, chains.AuthMiddleware)

	// 서버 구성
	http.Handle("/", chainedHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
