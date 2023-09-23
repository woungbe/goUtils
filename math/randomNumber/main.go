package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	MinAmt := 0.00001
	MaxAmt := 0.0001
	RanFloat(MinAmt, MaxAmt)

}

func RanFloat(min, max float64) {
	rand.Seed(time.Now().UnixNano())
	// 0.0001부터 0.001 사이의 랜덤한 부동 소수점 값을 생성합니다.
	// min := 0.0001
	// max := 0.001
	randomValue := min + rand.Float64()*(max-min)
	fmt.Printf("랜덤 값: %.6f\n", randomValue)
}

// 진짜 최소 수량 등록하기
