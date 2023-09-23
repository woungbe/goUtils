package main

import (
	"fmt"
	"strconv"
)

// type Account struct {
// 	Balance string
// }

type Account map[string]interface{}

func main() {
	// numbers := []Account{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	arrJson := []Account{
		{"Balance": "4"},
		{"Balance": "5"},
		{"Balance": "7"},
		{"Balance": "12"},
		{"Balance": "20"},
		{"Balance": "3"},
	}

	// FindMinMax(arrJson []Account, "Balance")
	max, min := FindMinMax(arrJson, "Balance")
	fmt.Println("max, min : ", max, min)

}

// max, min
func FindMinMax(arrJson []Account, key string) (int, int) {

	// min := math.MaxInt // 초기 최솟값을 int의 최댓값으로 설정
	var min int
	var max int
	for k, val := range arrJson {
		// bal, err := strconv.Atoi(val.Balance)
		tt := val[key].(string)
		bal, err := strconv.Atoi(tt)
		if err != nil {
			fmt.Println("err : ", err)
			return 0, 0
		}

		if k == 0 {
			min = bal
			max = bal
		} else {
			max = FindMax(max, bal)
			min = FindMin(min, bal)
		}
	}

	return max, min
}

func FindMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func FindMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
