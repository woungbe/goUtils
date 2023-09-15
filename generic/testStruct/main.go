package main

import (
	"encoding/json"
	"fmt"
)

// / test 1 ///
type Assets struct {
	AssetName string  `json:"assetName"`
	Value     float64 `json:"value"`
}

type REST_Accounst_RES[T any] struct {
	Success bool `json:"success"`
	Result  T    `json:"result"`
}

type REST_TEST struct {
	Name   string   `json:"name"`
	Assets []Assets `json:"assets"`
}

///// test 1 /////

/////// test 2 ////////

/////// test 2 ////////

func main() {

}

// 맞춤 제작하면 될것 같긴하네요? ..
func test1() {
	// 예제 데이터 생성
	assets := []Assets{
		{AssetName: "Bitcoin", Value: 45000.0},
		{AssetName: "Ethereum", Value: 3200.0},
	}

	res := REST_TEST{
		Name:   "HI",
		Assets: assets,
	}

	// REST_Accounst_RES 구조체 인스턴스 생성
	accounstRes := REST_Accounst_RES[REST_TEST]{
		Success: true,
		Result:  res,
	}

	// JSON 직렬화
	jsonData, err := json.Marshal(accounstRes)
	if err != nil {
		fmt.Println("JSON 직렬화 에러:", err)
		return
	}

	// JSON 출력
	fmt.Println(string(jsonData))
}

func test2() {

}
