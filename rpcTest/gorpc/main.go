package main

import (
	"fmt"

	rpcfunc "woungbe.utils/rpcTest/gorpc/rpcFunc"
)

/*
go rpc 작업 정리
설명: 서버에서 미리 지정한 function 으로 실행해서 reply 로 리턴 받는 방식

장점 : 빠르다.
단점 : 다른데서 사용하기 힘들 (그래서 GRPC가 있음 )


서버 - 클라이언트 끼리 struct 가 달라도 됨
다만 컬럼은 동일해야 받을 수 있음.

변경되거나 없으면 없는 부분만 누락됨!!
*/

func main() {

	_ = rpcfunc.GetPRCSvr()
	err := rpcfunc.StartRPCService()
	if err != nil {
		fmt.Println("err  : ", err)
	}

	done := make(chan bool)
	<-done

}
