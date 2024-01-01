package main

import (
	"fmt"
	"runtime"
)

/*
디버그 목적으로 사용 가능
프로그램 로직에서 동적 이름 사용하는 것은 권장하지 않음 ...


위 코드는 runtime.Caller와 runtime.FuncForPC 함수를 사용하여 현재 함수의 이름을 가져옵니다.
이것은 디버깅 또는 로깅 목적으로 사용될 수 있지만,
프로그램 로직에서 함수 이름을 동적으로 사용하는 것은 일반적으로 권장되지 않습니다.

*/

func myFunction() {
	pc, _, _, _ := runtime.Caller(0)
	functionName := runtime.FuncForPC(pc).Name()
	fmt.Println("현재 함수의 이름:", functionName)
}

func main() {
	myFunction()
}
