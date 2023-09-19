package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", CallbackHandler)
	http.ListenAndServe(":9910", nil)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// 요청 본문을 읽음
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Callback URL로 데이터 전달 (여기서는 간단히 출력)
	callbackURL := string(body)
	fmt.Printf("Received Callback URL: %s\n", callbackURL)

	// 요청에 대한 응답
	w.WriteHeader(http.StatusOK)
	response := "ok"
	fmt.Println(response)
	w.Write([]byte(response))
}
