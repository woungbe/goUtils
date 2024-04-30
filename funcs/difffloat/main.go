package main

import "fmt"

func main() {

	strAsk1Hoga := 0.1719
	strBid1Hoga := 0.1718

	price := 0.17190000

	if price <= strAsk1Hoga && price >= strBid1Hoga {
		fmt.Println("뭐라도 나와라?")
	}

}
