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
