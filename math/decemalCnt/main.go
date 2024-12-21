package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	price1 := "0.01010000"
	b, aa := DecimalCount(price1)
	if !b {
		fmt.Println("소수점 몇째라인지 파악할 수 없습니다.")
	}
	fmt.Println(aa)



	price2 := "0.00000001"
	c, cc := DecimalCount(price2)
	if !c {
		fmt.Println("소수점 몇째라인지 파악할 수 없습니다.")
	}
	fmt.Println(cc)


	price3 := "18.110000000000"
	i := GetFloatDotPos(price3)
	fmt.Println(i)

}

// 소수점 몇째자리인가 카운트 하는것
func DecimalCount(price string) (bool, int) {
	value, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return false, -1
	}

	var decimalPlaces int
	decimalPlaces = -1

	// 부동소수점 값을 문자열로 변환하여 소수점 이하 자릿수 계산
	strValue := strconv.FormatFloat(value, 'f', -1, 64)
	parts := strings.Split(strValue, ".")
	if len(parts) == 2 {
		decimalPlaces = len(parts[1])
	} else {
		if len(parts[0]) == 1 {
			i, _ := strconv.Atoi(parts[0])
			if i < 10 {
				return true, 0
			}
		}
		return false, decimalPlaces
	}
	return true, decimalPlaces
}


func GetFloatDotPos(fData string) int {
	if fData == "" {
		return 0
	}
	tmp := strings.Split(fData, ".")
	if len(tmp) > 1 {
		return len(tmp[1])
	}
	return 0
}