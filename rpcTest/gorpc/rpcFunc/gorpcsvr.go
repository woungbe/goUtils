package rpcfunc

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type RPCSvr struct {
}

type RPCFunc struct {
}

var RPCPTR *RPCSvr

func GetPRCSvr() *RPCSvr {
	if RPCPTR == nil {
		RPCPTR = new(RPCSvr)
	}
	return RPCPTR
}

// 서버 스타트
func StartRPCService() error {
	err := GetPRCSvr().Init()
	if err != nil {
		return err
	}
	go RPCPTR.startRPC()
	return nil
}

// 서버 시작
func (ty *RPCSvr) Init() error {
	return rpc.Register(new(RPCFunc))
}

func (ty *RPCSvr) startRPC() {

	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Error", "(ty *OKTAKycRPCSvr) StartRPC", err.Error())
		os.Exit(1)
		return
	}
	defer ln.Close()
	fmt.Println("RPC Receiver Start...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		defer conn.Close()

		go rpc.ServeConn(conn)
	}

	// fmt.Println("RPC Receiver Exit !!!")

}

func (ty *RPCFunc) RPCFuncSSS(args RPC_Func1_Reqeust, reply *RPC_Func1_Response) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Crit Error !!!!! Panic ", "Error", err)
		}
	}()

	reply.Send1 = args.Var1 + " 월드"
	reply.Send2 = args.Var1 + args.Var2 + " 월드2"
	reply.Send3 = "헬로월드"
	return nil
}
