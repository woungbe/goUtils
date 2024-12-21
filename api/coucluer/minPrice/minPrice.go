package main

import (
	"fmt"
)

// 각 회차의 투자 금액을 계산하는 함수
func calculateInvestmentRounds(totalInvestment, minInvestment, currentPrice, leverage float64, rounds int, multiplier float64) []float64 {
	investments := make([]float64, rounds)
	currentInvestment := minInvestment

	remainingInvestment := totalInvestment
	for i := 0; i < rounds; i++ {
		// 각 회차의 레버리지 적용 최소 투자 금액
		investment := currentInvestment * leverage

		// 현재가를 기준으로 최소 투자 금액을 만족하는 최소 수량 계산
		requiredQuantity := investment / currentPrice

		// 최소 수량을 만족하지 못할 경우, 최소 수량을 맞추기 위한 실제 투자 금액 계산
		if requiredQuantity < minInvestment/currentPrice {
			investment = minInvestment * leverage
		}

		// 남은 투자 금액보다 크면 남은 금액을 투자
		if investment > remainingInvestment*leverage {
			investment = remainingInvestment * leverage
		}

		investments[i] = investment / leverage
		remainingInvestment -= investment / leverage

		// 다음 회차를 위한 투자 금액 갱신
		currentInvestment *= multiplier
	}

	return investments
}

func main() {
	minInvestment := 5.0     // 최소 투자 금액 (증거금)
	currentPrice := 3830.0   // 현재가
	leverage := 20.0         // 레버리지
	rounds := 5              // 총 회차
	multiplier := 2.0        // 배율
	totalInvestment := 200.0 // 총 투자 금액

	investments := calculateInvestmentRounds(totalInvestment, minInvestment, currentPrice, leverage, rounds, multiplier)
	fmt.Printf("각 회차의 투자 금액 (증거금):\n")
	for i, investment := range investments {
		fmt.Printf("%d회차: %.2f USDT\n", i+1, investment)
	}

	sumInvestment := 0.0
	for _, investment := range investments {
		sumInvestment += investment
	}
	fmt.Printf("총 투자 증거금: %.2f USDT\n", sumInvestment)
}
