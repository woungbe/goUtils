package main

import (
	"fmt"
	"math"

	"woungbe.utils/covert/utils"
)

// USDT 계산 - Amount 계산을 해야지..
func FirstAmount(totalInvestment float64, rounds int, multiplier float64) []float64 {
	// 첫 회차의 투자 금액 계산
	firstInvestment := totalInvestment / ((1 - math.Pow(multiplier, float64(rounds))) / (1 - multiplier))

	// 각 회차의 투자 금액 계산
	investments := make([]float64, rounds)
	for i := 0; i < rounds; i++ {
		investments[i] = firstInvestment * math.Pow(multiplier, float64(i))
	}
	return investments
}

// 수량  = 금액 / 현재 가격  3830, 6.5, 0.001,
func Amount(price float64, usdt float64, leverage string, decimal string) string {
	// 수량 = 증거금 * 레버리지 / 가격   -- 소수점
	tprice, _ := utils.Float64(price)
	tusdt, _ := utils.Float64(usdt)
	tleverage, _ := utils.Float64(leverage)
	tdecimal, _ := utils.Float64(decimal)
	qty := tusdt * tleverage / tprice

	scaledValue := qty / tdecimal
	truncatedScaledValue := math.Trunc(scaledValue)
	finalValue := truncatedScaledValue * tdecimal
	send := utils.String(finalValue)
	return send
}

func main() {
	totalInvestment := 200.0 // 총 투자 금액
	rounds := 5              // 총 회차
	multiplier := 2.0        // 배율
	leverage := "20"         // 레버리지
	decimal := "1"           // thtnwja
	price := 0.0144790       // 1000 PEPEUSDT
	amountUSDT := FirstAmount(totalInvestment, rounds, multiplier)
	for _, v := range amountUSDT {
		qty := Amount(price, v, leverage, decimal)
		fmt.Println(price, qty)
	}
}
