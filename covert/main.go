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

}
