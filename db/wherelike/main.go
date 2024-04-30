package main

import (
	"fmt"
)

func main() {
	var coinName = "COE"
	var str = fmt.Sprintf("coin like '%%%s%%'", coinName)
	fmt.Println(str)
}
