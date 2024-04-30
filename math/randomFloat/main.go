package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 랜덤 실수 생성 함수
func randomFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func main() {
	// 랜덤 시드 초기화
	rand.Seed(time.Now().UnixNano())

	// 랜덤 실수 생성 및 출력
	randomNumber := randomFloat(0.0, 10.0)
	fmt.Printf("%.2f\n", randomNumber)
}
