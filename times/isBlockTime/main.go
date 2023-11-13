package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("서버 시작")

	for {
		b := IsBlockTime(06, 01)
		if b {
			// fmt.Printf("현재 한국 시간: %s, 데이터를 받지 않습니다.\n", koreaTime.Format("15:04:05"))
		} else {
			// fmt.Printf("현재 한국 시간: %s, 데이터를 받아서 처리합니다.\n", koreaTime.Format("15:04:05"))
		}

		time.Sleep(1 * time.Second)
	}

}

func IsBlockTime(start, end int) bool {
	loc, _ := time.LoadLocation("Asia/Seoul")
	koreaTime := time.Now().In(loc)

	// 현재 시간이 01:00부터 06:00 사이인지 확인합니다.
	if koreaTime.Hour() > end && koreaTime.Hour() < start {
		// 특정 시간대에는 데이터를 받지 않습니다.
		fmt.Printf("현재 한국 시간: %s, 데이터를 받지 않습니다.\n", koreaTime.Format("15:04:05"))
		return false
	} else {
		// 특정 시간대가 아니라면 데이터를 받아서 처리합니다.
		fmt.Printf("현재 한국 시간: %s, 데이터를 받아서 처리합니다.\n", koreaTime.Format("15:04:05"))
		return true
		// 여기에서 데이터를 받아서 처리하는 코드를 추가합니다.
		// 이 부분에 원하는 로직을 구현하세요.
	}
}
