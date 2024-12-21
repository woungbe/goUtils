package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ExchangeRateResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func main() {
	// API 호출 URL
	url := "https://open.er-api.com/v6/latest/USD"

	// HTTP GET 요청 보내기
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("API 요청 중 오류 발생: %v", err)
	}
	defer resp.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("응답 본문 읽기 중 오류 발생: %v", err)
	}

	// JSON 응답 파싱
	var exchangeRateResponse ExchangeRateResponse
	if err := json.Unmarshal(body, &exchangeRateResponse); err != nil {
		log.Fatalf("JSON 파싱 중 오류 발생: %v", err)
	}

	// KRW 환율 가져오기
	krwRate, exists := exchangeRateResponse.Rates["KRW"]
	if !exists {
		log.Fatalf("KRW 환율 정보를 찾을 수 없습니다.")
	}

	// KRW 환율 출력
	fmt.Printf("USD to KRW 환율: %f\n", krwRate)

}
