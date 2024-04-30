package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"woungbe.utils/websocket/clientconnect"
)

func main() {

	cnf := clientconnect.GetClientConnector() // 정리 하기
	cnf.Init("wss://dev-ccws.covest.pro/market")
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

			time.Sleep(time.Millisecond * 100)
		}
	}()

	<-done
}

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

	if strCMD == "send" {
		// { "Cmd":"TRI", "Symbol": "COEUSDT" }

		type CMDPacket struct {
			Cmd    string
			Symbol string
		}

		var packet CMDPacket
		packet.Cmd = "TRI"
		packet.Symbol = "COEUSDT"

		clientconnect.GetClientConnector().SendDataForQueue(packet)

	}
}
