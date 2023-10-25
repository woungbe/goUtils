package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	price := "0.01010000"
	// price := "10"
	b, aa := DecimalCount(price)
	if b == false {
		fmt.Println("소수점 몇째라인지 파악할 수 없습니다.")
	}
	fmt.Println(aa)

}

// 소수점 몇째자리인가 카운트 하는것
func DecimalCount(price string) (bool, int) {
	value, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("부동소수점 변환 오류:", err)
		return false, -1
	}

	var decimalPlaces int
	decimalPlaces = -1

	// 부동소수점 값을 문자열로 변환하여 소수점 이하 자릿수 계산
	strValue := strconv.FormatFloat(value, 'f', -1, 64)
	parts := strings.Split(strValue, ".")
	if len(parts) == 2 {
		decimalPlaces = len(parts[1])
		fmt.Println("소수점 이하 자릿수:", decimalPlaces)
	} else {
		if len(parts[0]) == 1 {
			i, _ := strconv.Atoi(parts[0])
			if i < 10 {
				return true, 0
			}
		}
		fmt.Println("부동소수점 형식이 아닙니다.")
		return false, decimalPlaces
	}

	return true, decimalPlaces
}
