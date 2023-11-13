package coboservice

import (
	"github.com/bitly/go-simplejson"
	"woungbe.utils/interface/custody/defined"
)

type IWebAPI interface {
	GetSupportCoin() (*simplejson.Json, *defined.ApiError) // 지원 코인 이름
	GetSupportedCoins(chainCode string) (*simplejson.Json, *defined.ApiError)
}
