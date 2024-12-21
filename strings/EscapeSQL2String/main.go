package main

import (
	"fmt"
	"strings"
)

type Info struct {
	Types     string
	Timestamp string
	Location  string
	Message   string
}

// escapeString은 문자열의 특수 문자를 이스케이프 처리합니다.
func escapeString(str string) string {
	// 문자열 내의 백슬래시와 작은따옴표를 이스케이프 처리
	return strings.ReplaceAll(strings.ReplaceAll(str, `\`, `\\`), `'`, `\'`)
}

func main() {
	inf := Info{
		Types:     "InternalCallBackError",
		Timestamp: "2024-06-26 10:55:23",
		Location:  "842262538",
		Message:   `{"code":-1021,"msg":"Timestamp for this request was 1000ms ahead of the server's time."} 400 : 842262538`,
	}

	// 입력 문자열을 이스케이프 처리
	escapedTypes := escapeString(inf.Types)
	escapedTimestamp := escapeString(inf.Timestamp)
	escapedLocation := escapeString(inf.Location)
	escapedMessage := escapeString(inf.Message)

	sql := fmt.Sprintf(`INSERT INTO FrontLog(types, reqeusttime, location, message) 
	VALUES ('%s', '%s', '%s', '%s')`, escapedTypes, escapedTimestamp, escapedLocation, escapedMessage)

	fmt.Println(sql)
}
