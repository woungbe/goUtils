package main

import (
	"fmt"
	"time"
)

func main() {
	// UTC 시간대 설정
	utcLocation := time.UTC

	now := time.Now().In(utcLocation)
	next := time.Date(now.Year(), now.Month(), now.Day(), 10, 11, 0, 0, utcLocation)

	if next.Sub(now) < 0 {
		fmt.Println("next가 더 작네")
		return
	}

	for {
		// 현재 시간(UTC 기준)
		now = time.Now()
		// 이미 오늘의 실행 시간이 지났다면, 다음 실행은 내일로 설정
		if now.After(next) {
			// next = next.Add(24 * time.Hour)
			next = next.Add(1 * time.Minute)
		}

		// 다음 실행까지 대기
		waitDuration := next.Sub(now)
		fmt.Printf("Next execution at %v (UTC)\n", next)
		time.Sleep(waitDuration)

		// 설정된 시간에 실행할 작업
		executeBatchJob()
	}
}

func executeBatchJob() {
	// 실제 배치 작업을 수행하는 코드
	fmt.Println("Executing batch job...")
	// 작업 완료 후 로그를 출력
	fmt.Println("Batch job completed at", time.Now().UTC(), "(UTC)")
}
