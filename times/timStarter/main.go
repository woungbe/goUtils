package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func startTimer(duration time.Duration, resetChan chan bool, doneChan chan bool) {
	startTime := time.Now() // 타이머 시작 시간 기록
	timer := time.NewTimer(duration)
	ticker := time.NewTicker(1 * time.Second) // 1초마다 틱하는 타이커 생성

	for {
		select {
		case <-timer.C:
			fmt.Println("타이머 완료, 함수 실행")
			doneChan <- true
			ticker.Stop() // 타이커 정지
			return
		case <-resetChan:
			if !timer.Stop() {
				<-timer.C
			}
			startTime = time.Now() // 리셋 시간 갱신
			fmt.Println("타이머 리셋됨")
			timer.Reset(duration)
		case <-ticker.C:
			elapsed := time.Since(startTime).Seconds()
			fmt.Printf("%.0f초 경과\n", elapsed)
		}
	}
}

var resetChan chan bool
var doneChan chan bool

func main() {
	resetChan = make(chan bool)
	doneChan = make(chan bool)

	go startTimer(10*time.Second, resetChan, doneChan)

	// 테스트를 위한 랜덤한 시간에서 타이머 리셋
	time.Sleep(2 * time.Second)
	resetChan <- true

	done := make(chan bool)
	r := bufio.NewReader(os.Stdin)
	go func() {
		for {
			line, err := r.ReadString('\n')
			if err != nil && err.Error() != "unexpected newline" {
				fmt.Println(err.Error())
				//	return
				line = ""
			}

			line = strings.TrimSpace(line)

			CMDPaser(line)

		}
	}()
	<-done
}

func CMDPaser(strCMD string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Crit Panic", "Error", err)
		}
	}()
	if strCMD == "" {
		return
	}
	if strCMD == "exit" {
		os.Exit(1)
		return
	}

	if strCMD == "start" {
		// resetChan <- true
		go startTimer(10*time.Second, resetChan, doneChan)
	}

	if strCMD == "send" {
		resetChan <- true
	}
}
