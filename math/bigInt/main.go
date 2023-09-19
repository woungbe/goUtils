package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {

	amount := ConvertAmount(0.01, 18)
	fmt.Println("amount : ", amount)

}

// 주문하기 -
func ConvertAmount(amount float64, pow int) *big.Int {
	// 0.01을 18자리 소수점 이하로 이동
	ethValue := new(big.Float).SetFloat64(amount)
	s := fmt.Sprintf("1e%d", pow)
	f, _ := strconv.ParseFloat(s, 10)
	ethValue.Mul(ethValue, big.NewFloat(f)) // 1e18은 10^18을 나타냅니다.

	// 소수점 이하를 정수로 변환
	ethInt := new(big.Int)
	ethValue.Int(ethInt)

	fmt.Println(ethInt)
	return ethInt
}
