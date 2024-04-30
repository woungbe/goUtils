package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// 근접한 키 찾기 함수
func findNearbyKeys(m map[string][]interface{}, target float64, count int) []string {
	// 맵의 키들을 배열에 저장
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	// 주어진 숫자에 가장 가까운 키 찾기
	sort.SliceStable(keys, func(i, j int) bool {
		ki, _ := strconv.ParseFloat(keys[i], 64)
		kj, _ := strconv.ParseFloat(keys[j], 64)
		return math.Abs(ki-target) < math.Abs(kj-target)
	})

	// 결과 배열 초기화
	var result []string

	// 주변 키 추가
	for i := 0; i < len(keys) && len(result) < count; i++ {
		result = append(result, keys[i])
	}

	return result
}

func main() {
	// 예시 맵
	myMap := make(map[string][]interface{})
	myMap["5.54"] = []interface{}{"a"}
	myMap["5.55"] = []interface{}{"b"}
	myMap["5.56"] = []interface{}{"c"}
	myMap["5.57"] = []interface{}{"d"}
	myMap["5.58"] = []interface{}{"e"}
	myMap["5.59"] = []interface{}{"f"}
	myMap["5.60"] = []interface{}{"g"}
	myMap["5.61"] = []interface{}{"h"}

	myMap["5.53"] = []interface{}{"h"}
	myMap["5.52"] = []interface{}{"h"}
	myMap["5.51"] = []interface{}{"h"}
	myMap["5.50"] = []interface{}{"h"}
	myMap["5.49"] = []interface{}{"h"}

	myMap["5.62"] = []interface{}{"h"}
	myMap["5.63"] = []interface{}{"h"}
	myMap["5.64"] = []interface{}{"h"}
	myMap["5.65"] = []interface{}{"h"}
	myMap["5.66"] = []interface{}{"h"}
	myMap["5.67"] = []interface{}{"h"}

	// 함수 사용
	nearbyKeys := findNearbyKeys(myMap, 5.57, 8)
	fmt.Println("근접한 키들:", nearbyKeys)
}
