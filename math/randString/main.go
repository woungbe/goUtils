package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 문자열 난수 생성 함수
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	randomStr := randomString(10) // 10자리 길이의 문자열 난수 생성
	fmt.Println(randomStr)
}
