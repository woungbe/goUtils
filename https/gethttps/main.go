package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 쿼리 파라미터 추출
	query := r.URL.Query()
	// name := query.Get("name") // 'name' 파라미터 값 추출
	symbol := query.Get("symbol")
	interval := query.Get("interval")
	startTime := query.Get("startTime")
	endTime := query.Get("endTime")
	limit := query.Get("limit")

	fmt.Println(symbol, interval, startTime, endTime, limit)

	if FindStringError(symbol, interval, startTime, endTime, limit) {
		fmt.Println("에러 엔딩")
		fmt.Fprintln(w, "에러 엔딩, ", symbol, interval, startTime, endTime, limit)
		return
	}

	// 응답으로 파라미터 값 반환
	fmt.Fprintln(w, "Hello, ", symbol, interval, startTime, endTime, limit)
}

func FindStringError(aa ...string) bool {
	// true : 에러,  false : 안에러
	for _, v := range aa {
		if v == "" || v == "0" {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("서버 시작!!")
	http.HandleFunc("/", handler)     // 루트 경로 핸들러 설정
	http.ListenAndServe(":8080", nil) // 서버 시작
}
