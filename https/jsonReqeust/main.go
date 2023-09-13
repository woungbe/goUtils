package main

import (
	"fmt"
)

type params map[string]interface{}

func main() {
	service()
}

func service() {
	m := make(params)
	m["subAccountId"] = 23423945
	m["txId"] = "ejklacnvkladjfkowef"
	tt := setJsonRequest(m)
	fmt.Println("tt : ", tt)
}

func setJsonRequest(m params) string {
	var val string
	val = `{`
	for k, v := range m {
		// r.setFormParam(k, v)
		switch v.(type) {
		case string:
			val += fmt.Sprintf(`"%s":"%s"`, k, v.(string))
		case int64:
			val += fmt.Sprintf(`"%s":%d`, k, v.(int64))
		case int:
			val += fmt.Sprintf(`"%s":%d`, k, v.(int))
		case float32:
			val += fmt.Sprintf(`"%s":%f`, k, v.(float32))
		case float64:
			val += fmt.Sprintf(`"%s":%f`, k, v.(float64))
		case bool:
			val += fmt.Sprintf(`"%s":%t`, k, v.(bool))
		case byte:
			val += fmt.Sprintf(`"%s":"%s"`, k, string(v.(byte)))
		}
	}

	val += `}`

	return val
}

// func main() {
// 	var x interface{} = "Hello" // 문자열 또는 int64 값을 가질 수 있는 빈 인터페이스

// 	switch value := x.(type) {
// 	case string:
// 		fmt.Printf("변수 x는 문자열입니다. 값: %s\n", value)
// 	case int64:
// 		fmt.Printf("변수 x는 int64입니다. 값: %d\n", value)
// 	default:
// 		fmt.Printf("변수 x는 다른 타입입니다. 타입: %s\n", reflect.TypeOf(x))
// 	}
// }
