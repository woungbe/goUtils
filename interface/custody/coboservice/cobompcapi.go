package coboservice

import (
	"fmt"

	"github.com/bitly/go-simplejson"
	"woungbe.utils/interface/custody/defined"
)

type CoboMPCAPI struct {
	Helloworld string
}

func GetCoboMPCAPI() *CoboMPCAPI {
	tt := new(CoboMPCAPI)
	return tt
}

func (ty *CoboMPCAPI) GetSupportCoin() (*simplejson.Json, *defined.ApiError) {

	fmt.Println("MPC API 처리 ")

	return nil, nil
}

func (ty *CoboMPCAPI) GetSupportedCoins(coin string) (*simplejson.Json, *defined.ApiError) {

	fmt.Println("MPC API 처리 ", coin)

	return nil, nil
}
