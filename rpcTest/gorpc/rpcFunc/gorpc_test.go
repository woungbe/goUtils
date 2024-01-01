package rpcfunc

import (
	"fmt"
	"testing"
)

func Default() {
	GETPRCClient().Init("127.0.0.1", "9000")
}

func TestRPCFunc1(t *testing.T) {
	Default()
	var in RPC_Func1_Reqeust
	reply := new(RPC_Func1_Response2)
	in.Var1 = "헬로"
	in.Var2 = "헬로2"

	err := GETPRCClient().RPCFunc1(in, reply)
	if err != nil {
		fmt.Println("err : ", err)
	}

	fmt.Printf("reply : %+v\n", reply)

}
