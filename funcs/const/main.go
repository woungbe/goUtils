package main

import "fmt"

var IntToName map[int]string

func init() {
	IntToName = make(map[int]string)
	IntToName[101] = "PENDING_APPROVAL"
	IntToName[201] = "QUEUED"
	IntToName[301] = "PENDING_SIGNATURE"
	IntToName[401] = "BROADCASTING"
	IntToName[403] = "PENDING_CONFIRMATION"
	IntToName[501] = "CONFIRMATION"
	IntToName[900] = "SUCCESS"
	IntToName[901] = "FAILED"

	// 다른 정수에 대한 매핑도 추가할 수 있습니다.
	// 예: IntToName[102] = "SOME_OTHER_CONSTANT"
}

func main() {
	// 101을 입력하면 해당하는 이름을 가져오는 함수 호출
	result := getIntName(101)
	fmt.Println("101에 대응하는 이름:", result)
}

// 정수 값을 이름에 매핑하는 맵을 이용하여 이름을 가져오는 함수
func getIntName(value int) string {
	name, exists := IntToName[value]
	if exists {
		return name
	}
	return "UNKNOWN" // 매핑이 없는 경우에는 "UNKNOWN"을 반환하거나 다른 처리를 할 수 있습니다.
}
