package main

import (
	"fmt"
	"runtime"
)

func logLineNumber() {
	pc, file, line, ok := runtime.Caller(1)
	fmt.Printf("%+v, %t\n", pc, ok)
	fmt.Printf("File: %s, Line: %d\n", file, line)
}

func main() {
	// 다음과 같이 원하는 위치에서 logLineNumber() 함수를 호출하여 현재 파일과 줄 번호를 로그로 출력할 수 있습니다.
	logLineNumber()
}
