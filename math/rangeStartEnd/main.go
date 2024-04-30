package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// 1~3번째 및 7번째 이상 키 찾기
func findSpecificKeys(m map[string][]interface{}, target float64, tickSize float64, lowerBound int, upperBound int) ([]string, []string) {
	// 맵의 키들을 배열에 저장
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	// 거리 계산을 위한 함수
	distance := func(k string) float64 {
		kv, _ := strconv.ParseFloat(k, 64)
		return math.Abs(kv - target)
	}

	// 키를 거리에 따라 정렬
	sort.Slice(keys, func(i, j int) bool {
		return distance(keys[i]) < distance(keys[j])
	})

	var lowerKeys []string // 1~3번째 키
	var upperKeys []string // 7번째 이상 키

	// 1~3번째 및 7번째 이상 키 추출
	for _, key := range keys {
		dist := distance(key)
		if dist < float64(lowerBound)*tickSize {
			if len(lowerKeys) < 3 {
				lowerKeys = append(lowerKeys, key)
			}
		} else if dist >= float64(upperBound)*tickSize {
			upperKeys = append(upperKeys, key)
		}
	}

	return lowerKeys, upperKeys
}

func main() {
	// 예시 맵
	myMap := make(map[string][]interface{})
	myMap["4.54"] = []interface{}{"a"}
	myMap["5.54"] = []interface{}{"a"}
	myMap["5.55"] = []interface{}{"b"}
	myMap["5.56"] = []interface{}{"c"}
	myMap["5.57"] = []interface{}{"d"}
	myMap["5.58"] = []interface{}{"e"}
	myMap["5.59"] = []interface{}{"f"}
	myMap["5.60"] = []interface{}{"g"}
	myMap["5.61"] = []interface{}{"h"}
	myMap["6.54"] = []interface{}{"a"}

	// 함수 사용
	lowerKeys, upperKeys := findSpecificKeys(myMap, 5.57, 0.01, 4, 7)
	fmt.Println("1~3번째 근접한 키들:", lowerKeys)
	fmt.Println("7번째 이상 근접한 키들:", upperKeys)
}
