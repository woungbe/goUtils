package main

import (
	"fmt"
	"time"
)

func main() {
	Unixtimes()
}

// 몇자리까지 나오는지 확ㅇ니하기
func Unixtimes() {

	fmt.Println("time.Now().Unix() : ", time.Now().Unix())
	fmt.Println("time.Now().UnixMilli() : ", time.Now().UnixMilli())
	fmt.Println("time.Now().UnixMicro() : ", time.Now().UnixMicro())

}
