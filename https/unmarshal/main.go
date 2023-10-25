package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func main() {
	// 예제 JSON 데이터
	jsonData := []byte(`{"name":"John"}`)

	// JSON 데이터를 구조체로 언마샬링
	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println("JSON 언마샬 오류:", err)
		return
	}

	// 언마샬된 데이터 출력
	fmt.Printf("이름: %s\n", person.Name)
	fmt.Printf("나이: %d\n", person.Age)
	fmt.Printf("이메일: %s\n", person.Email)
}
