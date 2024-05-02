package main

import (
	"log"
	"reflect"
	"runtime"
)

// 로그를 남기는 미들웨어 함수
func loggingMiddleware(fn func() int) func() int {
	return func() int {
		fnName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
		log.Printf("시작: %s", fnName)
		result := fn() // 실제 함수 실행
		log.Printf("종료: %s 결과: %d", fnName, result)
		return result
	}
}

// 예시 함수
func computeSomething() int {
	// 여기서 어떤 작업을 수행
	return 42 // 예시 결과
}

func main() {
	wrappedFunction := loggingMiddleware(computeSomething)
	result := wrappedFunction()
	log.Printf("함수 결과: %d", result)
}
