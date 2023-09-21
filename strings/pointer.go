package main

import (
	"encoding/json"
	"fmt"
)

type StringStruct struct {
	Str1 string
	Str2 string
	Str3 string
}

func main() {
	// null 또는 nil 문자열
	var str StringStruct
	var data = `{"Str1":"11", "Str2":null, "Str3":"33"}`

	err := json.Unmarshal([]byte(data), &str)
	if err != nil {
		fmt.Println("err : ", err)
	}

	// {11 0xc00002a120 33}
	// fmt.Printf("%+v", str)
	// fmt.Println("str.Str2 : ", str.Str2)

	fmt.Println(str, str.Str2)

}
