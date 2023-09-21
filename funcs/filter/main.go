package main

import "fmt"

type Balances struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

func main() {

	var balarray []Balances
	var send []Balances
	for _, val := range balarray {
		if val.Asset == "BTCUSDT" {
			send = append(send, val)
		}
	}

	fmt.Println(send)
}
