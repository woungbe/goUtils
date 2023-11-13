package coboservice

import (
	"fmt"

	"github.com/bitly/go-simplejson"
	"woungbe.utils/interface/custody/defined"
)

type CoboFullApi struct {
}

func (ty *CoboFullApi) Init() {
	fmt.Println("ㅋㅋㅋㅋㅋ")
}

func (ty *CoboFullApi) GetSupportCoin() (*simplejson.Json, *defined.ApiError) {

	fmt.Println("full custody API 처리 ")

	return nil, nil
}

func (ty *CoboFullApi) GetSupportedCoins(coin string) (*simplejson.Json, *defined.ApiError) {

	fmt.Println("full custody API 처리 ", coin)

	return nil, nil
}
