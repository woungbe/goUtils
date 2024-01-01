package smapleJson

import (
	"fmt"

	"github.com/bitly/go-simplejson"
)

func main() {


	data:map[chain_codes:[VET AVAXC ATP DOGE BTC ETC MATIC TOMO TRON ALGO HECO_HT ARBITRUM_ETH DOT LTC RBTC AURORA_ETH ETHW IOTX_IOTX BSC_BNB DASH ETH XZC LBTC_LBTC CMT FTM CRO OPT_ETH]]}

	jsonStr := `{"name": "John", "age": 30, "city": "New York"}`
	jsonData, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		fmt.Println("JSON 파싱 오류:", err)
		return
	}

	jsonData.Set("name", "Alice")
	jsonData.Set("age", 25)
	jsonData.Set("city", "Los Angeles")

	// 수정된 JSON 데이터를 문자열로 가져오기
	updatedJsonStr, err := jsonData.MarshalJSON()
	if err != nil {
		fmt.Println("JSON 직렬화 오류:", err)
		return
	}

	fmt.Println("Updated JSON 데이터:", string(updatedJsonStr))
}
