package main

import "fmt"

type SUCCESS struct {
	Values string
}

type ERROR struct {
	Success    bool // true, false
	Code       int
	Msg        string
	decription string
}

func main() {

	re, er := GetAccountDetails()
	if er.Success == false {
		fmt.Println("에러처리")
	}

	fmt.Printf("%+v", re)
}

// GetAccountDetails()  "/v1/custody/org_info/"
func GetAccountDetails() (SUCCESS, ERROR) {
	var req SUCCESS
	req.Values = "그래도 되나?"

	er := defaultErrorCode()

	er.Code = -120123
	er.Msg = "이러지는 말자."

	return req, er // 에러가 없는 경우, 두 번째 반환 값을 nil로 설정
}

func defaultErrorCode() ERROR {
	a := ERROR{Success: false}
	return a
}
