package main

import (
	"fmt"

	"woungbe.utils/funcs/attrFuncs/receive"
)

// 함수 타입 정의
type messageFuncType func(types, msg string)

// ClientConnector 구조체 정의
type ClientConnector struct {
	// 기존 필드...
	MessageFunc messageFuncType // 함수 타입 필드 추가
}

// 함수 타입을 사용하는 메소드 정의
func (c *ClientConnector) CallMessageFunc(types, aa string) {
	if c.MessageFunc != nil {
		c.MessageFunc(types, aa)
	} else {
		fmt.Println("messageFunc가 설정되지 않았습니다.")
	}
}

func ReceiveFunc(types, aa string) {
	if types == "V" {
		fmt.Println("브이")
	}
	if types == "인사" {
		fmt.Println("안녕하세요")
	}
}

func main() {
	// 함수 예제
	ex := receive.PackageEx{}

	// ClientConnector 인스턴스 생성 및 함수 할당
	connector := ClientConnector{
		MessageFunc: ReceiveFunc,
	}

	// 할당된 함수 호출
	connector.CallMessageFunc("V", "안녕하세요")
	connector.CallMessageFunc("인사", "안녕하세요")

	connector.MessageFunc = ex.Message

	connector.CallMessageFunc("V", "안녕하세요")
	connector.CallMessageFunc("인사", "안녕하세요")

}
