package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type URLInfo struct {
	URL   string `json:"url"`
	Types string `json:"types"`
}

func main() {

	// 파일에서 URL 읽기
	urls, err := ReadURLsFromFile("urls.json")
	if err != nil {
		fmt.Println(err)
	}

	// HTTP 핸들러 등록
	for _, url := range urls {
		http.HandleFunc(url.URL, FindHandler(url.Types))
	}

	// 서버 시작
	err = http.ListenAndServe(":9910", nil)
	if err != nil {
		log.Fatal("서버 시작 중 오류 발생:", err)
	}
}

// 파일에서 URL 읽어오는 함수
func ReadURLsFromFile(filename string) ([]URLInfo, error) {
	file, err := os.Open("urls.json")
	if err != nil {
		log.Fatal("파일을 열 수 없습니다:", err)
	}
	defer file.Close()

	var urls []URLInfo
	err = json.NewDecoder(file).Decode(&urls)
	if err != nil {
		log.Fatal("JSON 디코딩 중 오류 발생:", err)
	}
	return urls, nil
}

// HTTP 핸들러 콜백 함수
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func CallbackHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!2222")
}

func FindHandler(types string) func(http.ResponseWriter, *http.Request) {
	// if types == "json" {
	// 	return CallbackHandler
	// } else if types == "string" {
	// 	return CallbackHandler2
	// }
	return CallbackHandler
}
