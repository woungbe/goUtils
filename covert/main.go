package main

import (
	"fmt"
	"reflect"

	"woungbe.utils/covert/utils"
)

func main() {

	flot, _ := utils.Float64("123.123")
	fmt.Println("flot : ", reflect.TypeOf(flot))

	str := utils.String(0.5123)
	fmt.Println("utils.String : ", str, reflect.TypeOf(str))

	var i float32
	i = 120.123
	i2, err := utils.Int64(i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Int64 : ", i2, reflect.TypeOf(i2))

	secretEncKey := ""
	becData := []byte(secretEncKey)
	bytelen := utils.CheckByteLen(becData)
	fmt.Println("bytelen : ", bytelen)

}
