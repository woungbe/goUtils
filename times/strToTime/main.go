package main

import (
	"fmt"
	"time"
)

func main() {
	dateString := "2023-09-23 15:30:00"
	layout := "2006-01-02 15:04:05"

	// 문자열을 시간으로 변환
	resultTime, err := StringToTime(dateString, layout)
	if err != nil {
		fmt.Println("문자열을 시간으로 변환하는 데 실패했습니다:", err)
		return
	}

	// 1695483000
	// 1695483000000

	fmt.Println("변환된 시간:", resultTime.UnixMilli())
}

func StringToTime(input string, layout string) (time.Time, error) {
	parsedTime, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func TimeToString(input time.Time, layout string) string {
	return input.Format(layout)
}
