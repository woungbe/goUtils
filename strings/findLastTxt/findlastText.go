package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "BTCFDUSD1"
	if EndsWithFUSDT(str) {
		fmt.Println("있습니다.")
	} else {
		fmt.Println("없습니다.")
	}

}

func EndsWithFUSDT(input string) bool {
	// strings.HasSuffix 함수를 사용하여 문자열이 "FUSDT"로 끝나는지 확인합니다.
	return strings.HasSuffix(input, "FDUSD")
}
