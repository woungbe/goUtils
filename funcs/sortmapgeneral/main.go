package main

import (
	"fmt"
	"sort"
)

// Comparable 인터페이스는 키로 사용될 수 있는 타입에 대한 제약을 정의합니다.
type Comparable interface {
	~int | ~float64 | ~string
}

// SortMapForKey 함수는 제네릭을 사용하여 다양한 타입의 키를 정렬합니다.
func SortMapForKey[K Comparable, V any](map1 map[K]V) []K {
	var keys []K
	for k := range map1 {
		keys = append(keys, k)
	}

	// sort.Slice는 인터페이스를 사용하여 정렬을 수행합니다.
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}

func main() {
	// float64 키를 사용하는 맵 예제
	floatMap := map[float64]string{1.1: "a", 3.3: "c", 2.2: "b"}
	sortedFloatKeys := SortMapForKey(floatMap)
	fmt.Println("Sorted float64 keys:", sortedFloatKeys)

	// int 키를 사용하는 맵 예제
	intMap := map[int]string{10: "a", 30: "c", 20: "b"}
	sortedIntKeys := SortMapForKey(intMap)
	fmt.Println("Sorted int keys:", sortedIntKeys)

	// string 키를 사용하는 맵 예제
	stringMap := map[string]string{"apple": "a", "banana": "b", "cherry": "c"}
	sortedStringKeys := SortMapForKey(stringMap)
	fmt.Println("Sorted string keys:", sortedStringKeys)
}
