package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// JSON 데이터 준비 (맵 형식으로)
	jsonData := map[string]interface{}{
		"name":  "John Doe",
		"email": "johndoe@example.com",
	}

	// JSON 데이터를 바이트 슬라이스로 직렬화
	jsonDataBytes, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("JSON 직렬화 오류:", err)
		return
	}

	// 요청 생성
	url := "https://example.com/api"
	fmt.Println("어디로가지? : ", bytes.NewBuffer(jsonDataBytes))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonDataBytes))
	if err != nil {
		fmt.Println("HTTP 요청 생성 오류:", err)
		return
	}

	// 요청 헤더 설정 (JSON 형식임을 명시)
	req.Header.Set("Content-Type", "application/json")

	// HTTP 클라이언트 생성 및 요청 보내기
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP 요청 전송 오류:", err)
		return
	}
	defer resp.Body.Close()

	// 응답 처리
	fmt.Println("상태 코드:", resp.Status)
	// 여기서 추가적인 응답 처리를 수행할 수 있습니다.
}
