package main

import (
	"fmt"
	"strings"
)

// userIDX 값을 추출하여 "(1,2,3,4,5)" 형태의 문자열을 생성하는 함수
func FormatFromMapToInt(rowsMain map[int]map[string]interface{}, key string) (string, error) {
	var userIdxs []string

	for _, row := range rowsMain {
		if userIDX, ok := row[key]; ok {
			// userIDX 값이 int 형이라고 가정하고 문자열로 변환합니다.
			// 실제 타입에 따라 적절히 변환해야 할 수 있습니다.
			userIdxs = append(userIdxs, fmt.Sprintf("%v", userIDX))
		} else {
			// userIDX 키가 없는 경우 오류 처리
			return "", fmt.Errorf(fmt.Sprintf("%s key not found in row", key))
		}
	}

	// 슬라이스를 "(1,2,3,4,5)" 형태의 문자열로 변환
	result := fmt.Sprintf("(%s)", strings.Join(userIdxs, ","))

	return result, nil
}

// userIDX 값을 추출하여 "(1,2,3,4,5)" 형태의 문자열을 생성하는 함수
func FormatFromMapToString(rowsMain map[int]map[string]interface{}, key string) (string, error) {
	var userIdxs []string

	for _, row := range rowsMain {
		if userIDX, ok := row[key]; ok {
			// userIDX 값이 int 형이라고 가정하고 문자열로 변환합니다.
			// 실제 타입에 따라 적절히 변환해야 할 수 있습니다.
			userIdxs = append(userIdxs, fmt.Sprintf("'%v'", userIDX))
		} else {
			// userIDX 키가 없는 경우 오류 처리
			return "", fmt.Errorf(fmt.Sprintf("%s key not found in row", key))
		}
	}

	// 슬라이스를 "(1,2,3,4,5)" 형태의 문자열로 변환
	result := fmt.Sprintf("(%s)", strings.Join(userIdxs, ","))
	return result, nil
}

func main() {
	// 예제 데이터
	rowsMain := map[int]map[string]interface{}{
		0: {"userIDX": 1},
		1: {"userIDX": 2},
		2: {"userIDX": 3},
		3: {"userIDX": 4},
		4: {"userIDX": 5},
	}

	result, err := FormatFromMapToInt(rowsMain, "userIDX")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Formatted userIDX:", result)

	result1, err := FormatFromMapToString(rowsMain, "userIDX")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Formatted userIDX:", result1)

}
