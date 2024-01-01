package main

import (
	"fmt"
	"math/big"
)

func main() {

	NewBigInt("310", "18")

	NewBigInt2("310", 18)

}

func NewBigInt(amount string, decimal string) {

	amountBigInt := new(big.Int)
	amountBigInt.SetString(amount, 10)

	decimalBigInt := new(big.Int)
	decimalBigInt.SetString(decimal, 10)

	// amountBigInt를 10^decimalBigInt 만큼 곱하여 큰 정수를 생성합니다.
	result := new(big.Int).Mul(amountBigInt, new(big.Int).Exp(big.NewInt(10), decimalBigInt, nil))

	fmt.Println(result)
}

func NewBigInt2(amount string, decimal int) {

	amountBigInt := new(big.Int)
	amountBigInt.SetString(amount, 10)

	// 10^decimal을 계산합니다.
	exponent := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil)
	result := new(big.Int).Mul(amountBigInt, exponent)

	fmt.Println(result)

}
