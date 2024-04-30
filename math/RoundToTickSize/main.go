package main

import (
	"fmt"
	"math"
)

// RoundToTickSize 함수는 주어진 가격을 틱 사이즈의 배수로 반올림하고, 지정된 소수점 자릿수로 포맷합니다.
func RoundToTickSize(price float64, tickSize int, decimalPlaces float64) float64 {
	// 틱 사이즈에 따라 가격 반올림
	roundedPrice := math.Round(price/float64(tickSize)) * float64(tickSize)

	// 지정된 소수점 자릿수로 포맷
	shift := math.Pow(10, float64(decimalPlaces))
	return math.Round(roundedPrice*shift) / shift
}

func main() {
	// 예시: 가격을 틱 사이즈와 소수점 자릿수로 반올림
	price := 123.45678
	tickSize := 2
	decimalPlaces := 0.01

	roundedPrice := RoundToTickSize(price, tickSize, decimalPlaces)
	fmt.Printf("Rounded Price: %v\n", roundedPrice)
}
