package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 랜덤성을 위해 시드 설정
	rand.Seed(time.Now().UnixNano())

	// 문자열 슬라이스 정의
	items := []string{"Apple", "Banana", "Cherry", "Date", "Fig", "Grape"}

	// 슬라이스에서 무작위 항목 선택
	randomIndex := rand.Intn(len(items))
	randomItem := items[randomIndex]
	fmt.Println("무작위 항목:", randomItem)
}
