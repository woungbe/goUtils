package main

import (
	"fmt"
	"time"
)

type DaliyList struct {
	StartTime int64
	EndTime   int64
}

func main() {
	// 시작 날짜와 종료 날짜 설정
	// REST_startDate := "2023-09-01 00:00:00"
	// REST_endDate := "2023-09-11 00:00:00"
	// layout := "2006-01-02 15:04:05"
	// layoutDays := "2006-01-02"

	startDate := "2023-09-10"
	endDate := "2023-09-20"

	val := StartEndReturn(startDate, endDate)
	for k, v := range val {
		fmt.Println(k, " : ", v.StartTime, v.EndTime)
	}
}

func StartEndReturn(start, end string) []DaliyList {
	if len(start) <= 10 {
		start = start + " 00:00:00"
	}

	if len(end) <= 10 {
		end = end + " 00:00:00"
	}

	layout := "2006-01-02 15:04:05"
	startDate, err := stringToTime(start, layout)
	endDate, err := stringToTime(end, layout)
	if err != nil {
		fmt.Println("에러가 왜나유 ~ ")
	}

	// 하루씩 데이터를 호출할 반복문
	currentDate := startDate
	var send []DaliyList
	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		thistime := currentDate.UnixMilli()
		currentDate = currentDate.AddDate(0, 0, 1)

		send = append(send, DaliyList{StartTime: thistime, EndTime: currentDate.UnixMilli()})
	}

	return send
}

func stringToTime(input string, layout string) (time.Time, error) {
	parsedTime, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func timeToString(input time.Time, layout string) string {
	return input.Format(layout)
}
