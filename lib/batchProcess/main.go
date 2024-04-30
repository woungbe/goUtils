package main

import (
	"fmt"
	"time"
)

func main() {
	// 고루틴을 시작합니다.
	go func() {
		// 매 초마다 실행하기 위해 ticker 설정
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case t := <-ticker.C:
				// 현재 시간을 출력하고, 여기에 배치 작업을 실행할 함수를 호출
				fmt.Println("Batch job started at", t)
				executeBatchJob()
			}
		}
	}()

	// 메인 고루틴이 종료되지 않도록 대기합니다.
	// 실제로는 os.Signal 등을 사용하여 우아하게 종료할 수 있는 방법을 구현하는 것이 좋습니다.
	select {}
}

func executeBatchJob() {
	// 실제 배치 작업을 수행하는 코드
	fmt.Println("Executing batch job...")
	// 작업 완료 후 로그를 출력
	fmt.Println("Batch job completed.")
}
