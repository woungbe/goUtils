package rpcfunc

import (
	"fmt"
	"net/rpc"
)

type PRCClient struct {
	SvrIP string
	Port  string
}

var PRCClientPTR *PRCClient

func GETPRCClient() *PRCClient {
	if PRCClientPTR == nil {
		PRCClientPTR = new(PRCClient)
	}
	return PRCClientPTR
}

func (ty *PRCClient) Init(svrIP, port string) {
	ty.SvrIP = svrIP
	ty.Port = port
}

func (ty *PRCClient) Client() (*rpc.Client, error) {
	rpcAddr := fmt.Sprintf("%s:%s", ty.SvrIP, ty.Port)
	client, err2 := rpc.Dial("tcp", rpcAddr)
	if err2 != nil {
		if client != nil {
			client.Close()
		}
		return nil, err2
	}
	return client, nil
}

func (ty *PRCClient) RPCFunc1(in RPC_Func1_Reqeust, reply *RPC_Func1_Response2) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Crit Error !!!!! Panic ", "Error", err)
		}
	}()

	client, err := ty.Client()
	if err != nil {
		return err
	}

	rpcCall := client.Go("RPCFunc.RPCFuncSSS", in, reply, nil)
	<-rpcCall.Done
	client.Close()
	if rpcCall.Error != nil {
		return rpcCall.Error
	}
	return nil
}
