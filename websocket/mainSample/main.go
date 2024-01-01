package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"woungbe.utils/websocket/clientconnect"
)

func main() {

	cnf := clientconnect.GetClientConnector() // 정리 하기
	cnf.Init("ws://127.0.0.1:5105/")
	cnf.SvrConnect()

	done := make(chan bool)
	r := bufio.NewReader(os.Stdin)
	go func() {
		for {
			line, err := r.ReadString('\n')
			if err != nil && err.Error() != "unexpected newline" {
				fmt.Println(err.Error())
				//	return
				line = ""
			}

			line = strings.TrimSpace(line)

			CMDPaser(line)

		}
	}()
	<-done

}

// CMDPaser 콘솔 명령 파서
func CMDPaser(strCMD string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Crit Panic", "Error", err)
		}
	}()
	if strCMD == "" {
		return
	}
	if strCMD == "exit" {
		os.Exit(1)
		return
	}
	//명령을 공백을 기준으로 자른다.
	//cmds := strings.Split(strCMD, " ")
	if strCMD == "send" {

		// cnf := clientconnect.GetClientConnector() // 정리 하기

		// var packet COEDefines.WS_ADC_BalanceUpdate

		// packet.E = "BalanceUpdate"
		// packet.T = time.Now().UnixMilli()
		// packet.UX = 1033

		// packet.BalanceInfo.Symbol = symbol
		// packet.BalanceInfo.Balance = curBalance
		// packet.BalanceInfo.FreeBal = CurFreeBal
		// packet.BalanceInfo.LockedBal = CurLockedBal
		// packet.BalanceInfo.WithdrawLockedBal = CurWithdrawLockedBal

		// //==>웹소켓으로 유저 발란스 정보 변경 상태 전송 (쓰레드로 처리되도록 한다.)
		// GetADCClientConnector().SendDataForQueue(packet)

		// cnf.SendDataForQueue()

	}

}
