package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	// 클라이언트로부터 WebSocket 연결을 업그레이드합니다.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket 연결 업그레이드 실패:", err)
		return
	}
	defer conn.Close()

	fmt.Println("WebSocket 연결이 열렸습니다.")

	for {
		// 클라이언트로부터 메시지를 읽습니다.
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("메시지 수신 실패:", err)
			return
		}

		fmt.Printf("수신 메시지: %s\n", p)

		// 클라이언트로 메시지를 전송합니다.
		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println("메시지 전송 실패:", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html") // 웹 페이지 파일 (index.html)을 제공
	})
	http.HandleFunc("/ws", handleConnection) // WebSocket 연결을 처리할 핸들러

	fmt.Println("웹 서버가 8080 포트에서 실행 중...")
	http.ListenAndServe(":8080", nil)
}
