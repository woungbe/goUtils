package main

import (
	"fmt"
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	if objType.Kind() != reflect.Struct {
		panic("Input must be a struct")
	}

	result := make(map[string]interface{})

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		value := objValue.Field(i).Interface()
		result[field.Name] = value
	}

	return result
}

func MapToStringMap(input map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for key, value := range input {
		strValue := fmt.Sprintf("%v", value) // 값을 문자열로 변환
		result[key] = strValue
	}
	return result
}

func main() {
	// 예제로 사용할 구조체 생성
	type MyStruct struct {
		Name  string
		Age   int
		Email string
	}

	obj := MyStruct{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
	}

	// 구조체를 map[string]interface{}로 변환
	objMap := StructToMap(obj)

	// map[string]interface{}을 map[string]string으로 변환
	stringMap := MapToStringMap(objMap)

	// 결과 출력
	fmt.Println(stringMap)
}
