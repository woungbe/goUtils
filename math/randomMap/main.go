package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 맵에서 랜덤한 키들을 반환하는 함수
func getRandomKeys(m map[string]int, count int) []string {
	// 맵의 모든 키를 슬라이스에 저장
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// 랜덤 시드 설정
	rand.Seed(time.Now().UnixNano())

	// 반환할 키 슬라이스
	var randomKeys []string

	// 맵이 비어있거나 요청한 개수가 맵의 크기보다 크면, 모든 키를 반환
	if len(keys) == 0 || count >= len(keys) {
		return keys
	}

	// 랜덤 키 선택
	for len(randomKeys) < count {
		index := rand.Intn(len(keys))
		selectedKey := keys[index]

		// 중복 키 방지
		if !contains(randomKeys, selectedKey) {
			randomKeys = append(randomKeys, selectedKey)
		}
	}

	return randomKeys
}

// 키가 슬라이스에 포함되어 있는지 확인하는 함수
func contains(slice []string, key string) bool {
	for _, item := range slice {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	// 예시 맵
	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	// 맵에서 3개의 랜덤 키 추출
	randomKeys := getRandomKeys(myMap, 3)

	// 결과 출력
	fmt.Println("Random Keys:", randomKeys)
}
