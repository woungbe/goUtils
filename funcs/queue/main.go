package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func actions(element int) int {

	result := element * rand.Intn(10)
	return result
}

func processElement(element int, wg *sync.WaitGroup) {
	defer wg.Done() // 작업이 완료되면 WaitGroup 카운터 감소
	// 여기에서 요소를 가지고 작업을 수행하면 됩니다.
	result := actions(element)
	fmt.Printf("원본 요소: %d, 처리 결과: %d\n", element, result)
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	for _, element := range data {
		wg.Add(1) // WaitGroup 카운터 증가
		go processElement(element, &wg)
	}

	wg.Wait() // 모든 작업이 완료될 때까지 대기
}
