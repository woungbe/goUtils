package receive

import "fmt"

type PackageEx struct {
}

func (ty *PackageEx) Message(types, val string) {
	// 메시지 처리
	fmt.Println("Message가 올 수 있습니다.")

	if types == "V" {
		fmt.Println("브이")
	}
	if types == "인사" {
		fmt.Println("안녕하세요")
	}
}
