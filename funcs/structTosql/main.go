package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
	}

	insertQuery := GenerateInsertQuery("users", user)
	fmt.Println(insertQuery)
}

// GenerateInsertQuery 함수는 tableName으로 지정된 테이블에 data로 지정된 struct의 데이터를 INSERT하는 SQL 쿼리 문자열을 생성합니다.
func GenerateInsertQuery(tableName string, data interface{}) string {
	valueType := reflect.TypeOf(data)
	value := reflect.ValueOf(data)

	if valueType.Kind() != reflect.Struct {
		return ""
	}

	var columns []string
	var values []string

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fieldName := field.Name
		fieldValue := value.Field(i).Interface()
		// 값이 문자열이면 작은따옴표로 감싸줍니다.
		if strVal, ok := fieldValue.(string); ok {
			fieldValue = "'" + strVal + "'"
		}
		columns = append(columns, fieldName)
		values = append(values, fmt.Sprintf("%v", fieldValue))
	}

	columnsStr := strings.Join(columns, ", ")
	valuesStr := strings.Join(values, ", ")

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columnsStr, valuesStr)

	return query
}
