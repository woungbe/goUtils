package main

import (
	"encoding/json"
	"fmt"
)

type DataStruct struct {
	Data1 string `json:"data1"`
	Data2 string `json:"data2"`
	Data3 string `json:"data3"`
	Data4 string `json:"data4"`
}

type EventType struct {
	Evtypes string `json:"e"`
	//	EventTime int64  `json:"E"`
}

func main() {
	service()
}

func service() {
	data := `{"types":"1","data1":"123","data2":"123","data3":"123","data4":"123"}`
	var DataObj DataStruct

	err := json.Unmarshal([]byte(data), &DataObj)
	if err != nil {
		fmt.Println("err : ", err)
	}

	fmt.Println("DataObj : ", DataObj)
}

func service2() {
	/*
		이렇게 데이터가 들어오면. 비슷한 글자임으로 에러떨어짐.
		이럴때는 안쓰더라도. "E" 를 추가하면 에러가 안날 수 있음.
		** 아마 대문자 자동변환하면서 에러나는 걸로 추정 !!
	*/
	data := `{"e":"1","E":2}`
	var DataObj DataStruct

	err := json.Unmarshal([]byte(data), &DataObj)
	if err != nil {
		fmt.Println("err : ", err)
	}

	fmt.Println("DataObj : ", DataObj)
}
