package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 무작위 숫자들을 생성하고, 각 숫자의 빈도를 계산하는 함수
func generateRandomNumbers(count int) map[int]int {
	rand.Seed(time.Now().UnixNano())
	frequency := make(map[int]int)

	for i := 0; i < count; i++ {
		number := rand.Intn(10) + 1
		frequency[number]++
	}

	return frequency
}

// 가중치를 계산하는 함수
func calculateWeights(frequency map[int]int) map[int]float64 {
	weights := make(map[int]float64)
	maxFreq := 0
	for _, freq := range frequency {
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	for number, freq := range frequency {
		weights[number] = float64(maxFreq-freq) + 1 // 더 적게 나온 숫자에 더 높은 가중치 부여
	}

	return weights
}

func main() {
	frequency := generateRandomNumbers(100) // 예를 들어 100번 뽑기
	weights := calculateWeights(frequency)

	fmt.Println("빈도:", frequency)
	fmt.Println("가중치:", weights)
}
