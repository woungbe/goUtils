package main

import (
	"fmt"

	"github.com/bitly/go-simplejson"
)

func main() {
	jsonData := []byte(`{"name": "John", "age": 30, "city": "New York"}`)

	jsonObj, err := simplejson.NewJson(jsonData)
	if err != nil {
		panic(err)
	}

	// JSON 데이터 추출
	name := jsonObj.Get("name").MustString()
	age := jsonObj.Get("age").MustInt()
	city := jsonObj.Get("city").MustString()

	// 추출한 데이터로 struct 생성
	person := Person{Name: name, Age: age, City: city}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("City:", person.City)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}
