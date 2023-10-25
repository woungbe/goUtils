// Go에서 RPC (원격 프로시저 호출)를 구현하고 사용하는 예제를 제공합니다. RPC를 사용하면 네트워크를 통해 다른 프로세스 또는 컴퓨터에서 함수를 호출할 수 있습니다.

// 아래는 간단한 Go RPC 예제 코드입니다. 이 예제에서는 클라이언트와 서버 간에 간단한 덧셈 함수를 호출하도록 설정하겠습니다.

// 서버 코드 (server.go):
// go
// Copy code
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 덧셈 함수를 가진 객체
type Calculator int

// 덧셈 메서드
func (c *Calculator) Add(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// 덧셈 인자 구조체
type Args struct {
	A, B int
}

func main() {
	calculator := new(Calculator)
	rpc.Register(calculator)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("네트워크 리스닝 에러:", err)
		return
	}
	defer listener.Close()

	fmt.Println("RPC 서버 대기 중...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("연결 수락 에러:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
