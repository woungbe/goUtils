package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	upbitWebSocketURL := "wss://api.upbit.com/websocket/v1"

	// WebSocket 연결 설정
	conn, _, err := websocket.DefaultDialer.Dial(upbitWebSocketURL, nil)
	if err != nil {
		fmt.Println("WebSocket 연결에 실패했습니다:", err)
		return
	}
	defer conn.Close()

	// PING 메시지 전송 주기 설정
	pingInterval := 1 * time.Minute

	// PING 메시지를 주기적으로 보내고 PONG 응답을 처리하는 고루틴 실행
	go func() {
		for {
			// PING 메시지 보내기
			err := conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				fmt.Println("PING 메시지 전송 중 오류 발생:", err)
				return
			}
			fmt.Println("PING 메시지 전송 완료")

			// PONG 응답 대기
			conn.SetReadDeadline(time.Now().Add(10 * time.Second)) // 10초 내에 PONG을 기다립니다.
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("PONG 응답 대기 중 오류 발생:", err)
				return
			}
			if string(msg) == "PONG" {
				fmt.Println("PONG 응답 수신")
			}

			// 주기적으로 PING 메시지를 보내기 위해 대기
			time.Sleep(pingInterval)
		}
	}()

	// 채널해서 막으면 됨.
	done := make(chan bool)
	<-done

}
