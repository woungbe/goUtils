package main

import (
	"fmt"
	"time"
)

func main() {
	NextHour()
}

func NextHour() {
	// 현재 시간을 얻어옵니다.
	now := time.Now()
	fmt.Println("현재 시간:", now)

	// 다음 정각을 계산합니다.
	nextHour := now.Add(time.Hour)
	nextHour = time.Date(nextHour.Year(), nextHour.Month(), nextHour.Day(), nextHour.Hour(), 0, 0, 0, nextHour.Location())
	fmt.Println("다음 정각:", nextHour)

	// 다음 정각까지 대기합니다.
	time.Sleep(nextHour.Sub(now))

	// 정각에 실행할 동작을 여기에 추가합니다.
	fmt.Println("정각입니다! 동작을 실행합니다.")
}
