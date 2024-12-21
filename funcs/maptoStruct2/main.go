package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// MapToStruct는 맵 데이터를 구조체로 복사합니다.
func MapToStruct(tag string, mapData map[string]interface{}, result interface{}) error {
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
		fieldValue := resultValue.Elem().Field(i)

		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			// 중첩된 구조체의 필드를 상위 구조체에 병합하여 처리
			err := MapToStruct(tag, mapData, fieldValue.Addr().Interface())
			if err != nil {
				return err
			}
			continue
		}

		fieldName := field.Tag.Get(tag)
		mapValue, exists := mapData[fieldName]
		if !exists {
			continue // 맵에 필드 이름이 없는 경우 스킵
		}

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

	if mapValueType == nil {
		if fieldType.String() == "string" {
			fieldValue.SetString("")
			return nil
		}

		// 문자열(string)을 int64로 변환해서 할당하는 로직 추가
		if fieldType.String() == "int64" {
			fieldValue.SetInt(0)
			return nil
		}

		// bool
		if fieldType.String() == "bool" {
			fieldValue.SetBool(false)
			return nil
		}

		// float64
		if fieldType.String() == "float64" {
			fieldValue.SetFloat(0)
			return nil
		}

	} else {
		// 맵 값 타입을 구조체 필드 타입으로 변환 가능한지 확인
		if mapValueType.ConvertibleTo(fieldType) {
			*fieldValue = reflect.ValueOf(mapValue).Convert(fieldType)
			return nil
		}

		if fieldType.Kind() == reflect.Int {
			strValue := mapValue.(string)
			val, err := strconv.Atoi(strValue)
			if err != nil {
				return err
			}
			fieldValue.SetInt(int64(val))
			return nil
		}

		// 문자열(string)을 int64로 변환해서 할당하는 로직 추가
		if fieldType.Kind() == reflect.Int64 {
			strValue := mapValue.(string)
			intValue, err := strconv.ParseInt(strValue, 10, 64)
			if err != nil {
				return err
			}
			fieldValue.SetInt(intValue)
			return nil
		}

		// bool
		if fieldType.Kind() == reflect.Bool {
			var send bool
			boolValue := mapValue.(string)

			tmp, err := strconv.Atoi(boolValue)
			if err != nil {
				return err
			}

			if tmp == 0 {
				send = false
			} else {
				send = true
			}
			fieldValue.SetBool(send)
			return nil
		}

		// float64
		if fieldType.Kind() == reflect.Float64 {
			strValue := mapValue.(string)
			floatval, err := strconv.ParseFloat(strValue, 64)
			if err != nil {
				return err
			}
			fieldValue.SetFloat(floatval)
			return nil
		}
	}

	// 나머지 타입 변환 로직 추가

	return fmt.Errorf("타입 변환 불가능: %v → %v", mapValueType, fieldType)
}

// 테스트용 구조체들
type AA struct {
	Name string `json:"name"`
}

type BB struct {
	AA
	Age int `json:"age"`
}

func main() {
	mapData := map[string]interface{}{
		"name": "John",
		"age":  "30",
	}

	var bb BB
	err := MapToStruct("json", mapData, &bb)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		// %+v 형식으로 출력하여 중첩된 구조체의 필드가 병합된 형태로 표시되도록 합니다.
		fmt.Printf("Result: %+v\n", bb)
	}
}
