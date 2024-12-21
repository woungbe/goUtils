package main

import (
	"encoding/json"
	"fmt"

	"woungbe.utils/https/wsSample"
)

type CMDPacket struct {
	Cmd    string `json:"cmd"`
	Symbol string `json:"symbol"`
}

func cbOnConnect() {
	fmt.Println("cbOnConnect")
}
func cbOnUnConnect() {
	fmt.Println("cbOnUnConnect")
}
func cbOnMessage(msg []byte) {
	fmt.Println("cbOnMessage", string(msg))
}

func main() {

	wsObj := wsSample.BinanceUserWSObject{}
	wsObj.Init("wss://ccws.ggex.io/market", "")
	wsObj.SetCallbackFunc(cbOnConnect, cbOnUnConnect, cbOnMessage)
	wsObj.ClientConnect()

	var packet CMDPacket
	packet.Cmd = "TRI"
	packet.Symbol = "RPGUSDT"
	jsonData, er := json.Marshal(packet)
	if er != nil {
		fmt.Println("json.Marshal error", er)
		return
	}
	wsObj.SendMessage(jsonData) 

	select {} // 프로그램이 종료되지 않도록 무한 대기

}