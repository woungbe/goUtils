package main

/*
	보통 struct 할때 unMarshal 해서 짜증났다면.
	그걸 function 화 시키는 방법임

	아래는 그닥 특이할게 없는데.
	>> SetData(must []byte, structs interface{})

	중요한건 넣는 방법임.
	*** 넣을 & 주소로 넣어야됨 !!
	ty.SetData([]byte(data), &resultA)

	& : 주소를 가르킴
	* : 주소일때 변수를 가르킴

	a :=42
	b := &a  // print..  0xadf13123
	c := *b // print // 42

*/

import (
	"encoding/json"
	"fmt"
)

type CoboCtrl struct {
}

func (ty *CoboCtrl) SetData(must []byte, structs interface{}) {
	er2 := json.Unmarshal(must, structs)
	if er2 != nil {
		fmt.Println("utils.MapToStruct : ", er2)
	}
}

type StructA struct {
	Field1 string `json:"field1"`
}

type StructB struct {
	Field2 int `json:"field2"`
}

func main() {
	data := `{"field1": "Hello"}` // JSON 데이터 예시

	var resultA StructA
	var resultB StructB

	var ty CoboCtrl

	ty.SetData([]byte(data), &resultA)
	fmt.Println("Result from StructA:", resultA)

	ty.SetData([]byte(data), &resultB)
	fmt.Println("Result from StructB:", resultB)
}
