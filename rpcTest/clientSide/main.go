package main

import (
	"fmt"
	"net/rpc"
)

// 덧셈 인자 구조체
type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("RPC 클라이언트 연결 에러:", err)
		return
	}
	defer client.Close()

	args := Args{5, 3}
	var reply int

	err = client.Call("Calculator.Add", args, &reply)
	if err != nil {
		fmt.Println("RPC 호출 에러:", err)
		return
	}

	fmt.Printf("덧셈 결과: %d + %d = %d\n", args.A, args.B, reply)
}
