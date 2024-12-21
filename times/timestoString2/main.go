package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// 문자열로 주어진 타임스탬프
	timestampStr := "1705401833000"

	// 문자열을 정수형으로 변환
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
		return
	}

	// 밀리초를 나노초로 변환
	t := time.Unix(0, timestamp*int64(time.Millisecond))

	// 원하는 형식으로 날짜 출력
	formattedTime := t.Format("2006-01-02 15:04:05")
	fmt.Println(formattedTime)
}
