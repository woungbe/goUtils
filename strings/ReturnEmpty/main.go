package main

import "fmt"

func main() {

	val := "체크"
	val2 := "체크2"
	val3 := "체크3"
	returnVal := " value like '%%%s%%' "
	str := ReturnEmptyStr(returnVal, val, val2, val3)

	fmt.Println(str)

}

//
func ReturnEmptyStr(returnVal string, val ...string) string {
	send := ""
	if len(val) > 0 {
		// []string을 []any로 변환
		anyVal := make([]any, len(val))
		for i, v := range val {
			anyVal[i] = v
		}
		send = fmt.Sprintf(returnVal, anyVal...)
	}

	return send
}
