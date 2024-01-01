package main

import "fmt"

func main() {

	if FindStringError("aa", "bb", "cc", "dd", "ff", "gg", "hh", "") {
		fmt.Println("에러다")
	} else {
		fmt.Println("안 에러 났다.")
	}

}

func FindStringError(aa ...string) bool {
	// true : 에러,  false : 안에러
	for _, v := range aa {
		if v == "" || v == "0" {
			return true
		}
	}

	return false
}
