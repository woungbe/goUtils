package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// MapToStruct 함수는 맵을 구조체로 변환하는 함수입니다.
// mapData: 변환할 맵 데이터
// result: 변환된 구조체 결과
func MapToStruct(mapData map[string]interface{}, result interface{}) error {
	// 입력된 결과 인터페이스의 유효성을 검사
	resultValue := reflect.ValueOf(result)
	if resultValue.Kind() != reflect.Ptr || resultValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("result 인자는 구조체 포인터여야 합니다")
	}

	// 결과 구조체의 타입 정보를 가져옴
	resultType := resultValue.Elem().Type()

	// 맵 데이터를 구조체로 복사
	for i := 0; i < resultType.NumField(); i++ {
		field := resultType.Field(i)
		fieldName := field.Tag.Get("json")
		mapValue, exists := mapData[fieldName]
		if !exists {
			continue // 맵에 필드 이름이 없는 경우 스킵
		}

		fieldValue := resultValue.Elem().Field(i)

		if fieldValue.Type() == reflect.TypeOf(mapValue) {
			mapValueReflect := reflect.ValueOf(mapValue)
			fieldValue.Set(mapValueReflect.Convert(fieldValue.Type()))
			continue
		}

		// 맵 값을 구조체 필드로 할당 (타입 변환)
		if err := assignValue(mapValue, &fieldValue); err != nil {
			return fmt.Errorf("필드 %s: %v", fieldName, err)
		}
	}

	return nil
}

// assignValue 함수는 맵 값을 구조체 필드로 할당합니다.
func assignValue(mapValue interface{}, fieldValue *reflect.Value) error {
	// 구조체 필드 타입과 맵 값 타입 확인
	fieldType := fieldValue.Type()
	mapValueType := reflect.TypeOf(mapValue)

	// 맵 값 타입을 구조체 필드 타입으로 변환 가능한지 확인
	if mapValueType.ConvertibleTo(fieldType) {
		*fieldValue = reflect.ValueOf(mapValue).Convert(fieldType)
		return nil
	}

	// 문자열(string)을 int64로 변환해서 할당하는 로직 추가
	if fieldType.Kind() == reflect.Int64 && mapValueType.Kind() == reflect.String {
		strValue := mapValue.(string)
		intValue, err := strconv.ParseInt(strValue, 10, 64)
		if err != nil {
			return err
		}
		fieldValue.SetInt(intValue)
		return nil
	}

	// 나머지 타입 변환 로직 추가 (float 등)

	return fmt.Errorf("타입 변환 불가능: %v → %v", mapValueType, fieldType)
}

// 구조체 정의
type MyStruct struct {
	Field1 string  `json:"Field1"`
	Field2 int64   `json:"Field2"`
	Field3 float64 `json:"Field3"`
}

func main() {
	// 예제 맵 데이터
	mapData := map[string]interface{}{
		"Field1": "Hello",
		"Field2": "42", // 문자열로 오는 데이터
		"Field3": 3.14159,
	}

	// 변환할 구조체 생성
	var myStruct MyStruct

	// MapToStruct 함수를 사용하여 맵을 구조체로 변환
	if err := MapToStruct(mapData, &myStruct); err != nil {
		fmt.Println("오류:", err)
		return
	}

	// 변환된 구조체 사용
	fmt.Println("Field1:", myStruct.Field1)
	fmt.Println("Field2:", myStruct.Field2)
	fmt.Println("Field3:", myStruct.Field3)
}
