package main

import (
	"fmt"
	"math"
)

func main() {
	value := 3.14159265
	lotsize := 4
	roundedValue := RoundFloat(value, lotsize)
	fmt.Println("roundedValue : ", roundedValue)
	//fmt.Printf("원래 값: %f\n", value)
	//fmt.Printf("반올림된 값 (%d 자리까지): %.3f\n", lotsize, roundedValue)
	// := fmt.Sprintf("반올림된 값 (%d 자리까지): %.3f\n", lotsize, roundedValue);
}

func RoundFloat(num float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	rounded := math.Round(num*shift) / shift
	return rounded
}
