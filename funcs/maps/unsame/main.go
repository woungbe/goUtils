package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonData := data() // 주어진 JSON 데이터를 여기에 넣으세요.

	// JSON 데이터를 map[string]interface{}로 언마샬링
	var data []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		fmt.Println("JSON 데이터 파싱 오류:", err)
		return
	}

	for _, item := range data {
		filterType, ok := item["filterType"].(string)
		if !ok {
			fmt.Println("filterType 필드가 문자열이 아닙니다.")
			continue
		}

		if filterType == "PRICE_FILTER" {
			minPrice, b1 := item["minPrice"].(string)
			if !b1 {
				fmt.Println("minPrice 필드가 문자열이 아닙니다.")
				continue
			}

			maxPrice, b1 := item["maxPrice"].(string)
			if !b1 {
				fmt.Println("maxPrice 필드가 문자열이 아닙니다.")
				continue
			}

			tickSize, b1 := item["tickSize"].(string)
			if !b1 {
				fmt.Println("tickSize 필드가 문자열이 아닙니다.")
				continue
			}
			fmt.Println(minPrice, maxPrice, tickSize)
		}

		fmt.Println("Filter Type:", filterType)
	}

}

func data() string {
	return `[
        {
            "filterType": "PRICE_FILTER",
            "minPrice": "0.01000000",
            "maxPrice": "1000000.00000000",
            "tickSize": "0.01000000"
        },
        {
            "filterType": "LOT_SIZE",
            "minQty": "0.00001000",
            "maxQty": "9000.00000000",
            "stepSize": "0.00001000"
        },
        {
            "filterType": "ICEBERG_PARTS",
            "limit": 10
        },
        {
            "filterType": "MARKET_LOT_SIZE",
            "minQty": "0.00000000",
            "maxQty": "143.62962720",
            "stepSize": "0.00000000"
        },
        {
            "filterType": "TRAILING_DELTA",
            "minTrailingAboveDelta": 10,
            "maxTrailingAboveDelta": 2000,
            "minTrailingBelowDelta": 10,
            "maxTrailingBelowDelta": 2000
        },
        {
            "filterType": "PERCENT_PRICE_BY_SIDE",
            "bidMultiplierUp": "5",
            "bidMultiplierDown": "0.2",
            "askMultiplierUp": "5",
            "askMultiplierDown": "0.2",
            "avgPriceMins": 5
        },
        {
            "filterType": "NOTIONAL",
            "minNotional": "5.00000000",
            "applyMinToMarket": true,
            "maxNotional": "9000000.00000000",
            "applyMaxToMarket": false,
            "avgPriceMins": 5
        },
        {
            "filterType": "MAX_NUM_ORDERS",
            "maxNumOrders": 200
        },
        {
            "filterType": "MAX_NUM_ALGO_ORDERS",
            "maxNumAlgoOrders": 5
        }
    ]`
}
