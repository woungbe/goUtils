package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New() // 새로운 cron 스케줄러 생성

	// 매 분마다 실행되는 작업 추가
	_, err := c.AddFunc("* * * * *", func() {
		fmt.Println("Every minute task", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		fmt.Println("Error scheduling task:", err)
		return
	}

	// 매일 12시 30분에 실행되는 작업 추가
	_, err = c.AddFunc("30 12 * * *", func() {
		fmt.Println("Daily task at 12:30", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		fmt.Println("Error scheduling task:", err)
		return
	}

	// 스케줄러 시작
	c.Start()

	// 주 프로그램이 종료되지 않도록 대기
	select {}
}
