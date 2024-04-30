package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	난수 발생기 2.20 버전임

*/

func main() {
	// seed := int64(12345) // 고정된 시드 값
	seed := int64(time.Now().UnixNano())
	src := rand.NewSource(seed)
	rng := rand.New(src) // 개별 난수 생성기

	randomNumber := rng.() // 난수 생성
	fmt.Println(randomNumber)
}
