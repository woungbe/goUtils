package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	aa := randomString(32)
	bb := randomInt(14)

	fmt.Println(aa)
	fmt.Println(bb)

}

func randomString(length int) string {
	// 무작위 문자들의 집합
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 무작위 문자열 생성
	randomBytes := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range randomBytes {
		randomBytes[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomBytes)
}

func randomInt(length int) string {
	// 무작위 문자들의 집합
	const charset0 = "123456789"
	const charset = "0123456789"

	// 무작위 문자열 생성
	randomBytes := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range randomBytes {
		if i == 0 {
			randomBytes[i] = charset0[rand.Intn(len(charset0))]
		} else {
			randomBytes[i] = charset[rand.Intn(len(charset))]
		}
	}
	return string(randomBytes)
}
