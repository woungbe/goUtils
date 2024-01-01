package main

import (
	"fmt"
	"net/rpc"
)

type RPC_OKTA_KYC_Status struct {
	KycType            string //KYC type = USER / ENTERPRISE
	BrokerSubAccountId string //sub account id
}

type RPC_OKTA_KYC_Status_Resp struct {
	ErrorCode  string //에러코드  = 	000000000 =성공 그외 오류
	StatusInfo string
	LevelInfo  string
	NonExisted string
}

// 사용자 KYC URL 생성
type RPC_OKTA_KYC_GenerateURL struct {
	KycType            string //KYC type = USER / ENTERPRISE
	UserIDX            int64  //유저IDX
	BrokerSubAccountId string //sub account id
	RedirectUrl        string //kyc verification flow end redirect	url
}

type RPC_OKTA_KYC_GenerateURL_Resp struct {
	ErrorCode  string //에러코드  = 	000000000 =성공 그외 오류
	PageUrl    string // kyc	verification	url
	KycTransId string // kyc	transaction	id
	Timestamp  int64
}

func main() {
	var svrIP string = "127.0.0.1"
	var args RPC_OKTA_KYC_Status
	args.KycType = "USER"
	args.BrokerSubAccountId = "2885218267026853632"
	reply := new(RPC_OKTA_KYC_Status_Resp)
	err := RPCClient_OKTA_KYC_Status(svrIP, args, reply)
	if err != nil {
		fmt.Println("err : ", err)
	}

	fmt.Printf("reply : %+v\n", reply)
	/*
		var in RPC_OKTA_KYC_GenerateURL
		reply1 := new(RPC_OKTA_KYC_GenerateURL_Resp)

		err = RPCClient_OKTA_KYC_GenerateURL(svrIP, in, reply1)
		if err != nil {
			fmt.Println("err : ", err)
		}
	*/

}

func RPCClient_OKTA_KYC_Status(svrIP string, args RPC_OKTA_KYC_Status, reply *RPC_OKTA_KYC_Status_Resp) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Crit Error !!!!! Panic ", "Error", err)
		}
	}()

	rpcAddr := fmt.Sprintf("%s:10911", svrIP)
	client, err2 := rpc.Dial("tcp", rpcAddr)
	fmt.Println("client : ", client)
	if err2 != nil {
		if client != nil {
			client.Close()
		}
		fmt.Println("Error", "RPC server not connect", err2.Error())
		return err2
	}
	rpcCall := client.Go("OKTAKycRPCFunc.RPCFunc_OKTA_KYC_Status", args, reply, nil)
	<-rpcCall.Done
	client.Close()
	if rpcCall.Error != nil {
		return rpcCall.Error
	}
	return nil
}

func RPCClient_OKTA_KYC_GenerateURL(svrIP string, args RPC_OKTA_KYC_GenerateURL, reply *RPC_OKTA_KYC_GenerateURL_Resp) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Crit Error !!!!! Panic ", "Error", err)
		}
	}()

	rpcAddr := fmt.Sprintf("%s:10911", svrIP)
	client, err2 := rpc.Dial("tcp", rpcAddr)
	if err2 != nil {
		if client != nil {
			client.Close()
		}
		fmt.Println("Error", "RPC server not connect", err2.Error())
		return err2
	}
	rpcCall := client.Go("OKTAKycRPCFunc.RPCFunc_OKTA_KYC_GenerateURL", args, reply, nil)
	<-rpcCall.Done
	client.Close()
	if rpcCall.Error != nil {
		return rpcCall.Error
	}
	return nil
}
