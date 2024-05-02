package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("시작은 했습니다.")
	// OS 종료 신호를 받기 위한 채널 설정
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 신호를 기다림
	select {
	case sig := <-sigs:
		// 종료 신호를 받았을 때 처리
		println("받은 신호:", sig)
	}
}
