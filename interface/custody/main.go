package main

import (
	"woungbe.utils/interface/custody/coboservice"
)

func main() {
	/*
		//  * 붙으면 & ,   안붙어 있으면 처리하기
		(ty CoboMPCAPI)  , 	webapi = coboservice.CoboFullApi{}
		(ty *CoboMPCAPI) , 	webapi = &coboservice.CoboFullApi{}
	*/

	var webapi coboservice.IWebAPI
	fullapi := new(coboservice.CoboFullApi)
	fullapi.Init()
	webapi = fullapi
	webapi.GetSupportCoin()
	webapi.GetSupportedCoins("BTCUSDT")

	webapi = &coboservice.CoboMPCAPI{}
	webapi.GetSupportCoin()
	webapi.GetSupportedCoins("BTCUSDT")

}
