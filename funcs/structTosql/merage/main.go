package main

import (
	"fmt"
	"reflect"
	"strings"
)

// 없으면 insert , 있으면 update

type DepositMaster struct {
	Txid                string `json:"txid" db:"txid"`
	CoboID              string `json:"coboID" db:"coboID"`
	Status              string `json:"status" db:"status"`
	Coin                string `json:"coin" db:"coin"`
	Network             string `json:"network" db:"network"`
	FromAddress         string `json:"fromAddress" db:"fromAddress"`
	ToAddress           string `json:"toAddress" db:"toAddress"`
	ConfirmedNumber     int    `json:"confirmedNumber" db:"confirmedNumber"`
	ConfirmingThreshold int    `json:"confirmingThreshold" db:"confirmingThreshold"`
	TransactionType     int    `json:"transactionType" db:"transactionType"`
	CreatedTime         int    `json:"createdTime" db:"createdTime"`
	UpdatedTime         int    `json:"updatedTime" db:"updatedTime"`
	GasPrice            int    `json:"gasPrice" db:"gasPrice"`
	GasLimit            int    `json:"gasLimit" db:"gasLimit"`
	FeeUsed             int    `json:"feeUsed" db:"feeUsed"`
}

func main() {
	user := DepositMaster{
		Txid:                "Txid",
		CoboID:              "CoboID",
		Status:              "500",
		Coin:                "Coin",
		Network:             "Network",
		FromAddress:         "0x123123",
		ToAddress:           "0x123124",
		ConfirmedNumber:     7,
		ConfirmingThreshold: 64,
		TransactionType:     900,
		CreatedTime:         1613912309123,
		UpdatedTime:         16192381923891,
		GasPrice:            1000000,
		GasLimit:            1000000,
		FeeUsed:             200000000,
	}

	// query := GenerateInsertOrUpdateQuery("users", user)
	// /// query := GenerateInsertOrUpdateQuery("users", user, "Username = VALUES(Username)")
	query := GenerateInsertOrUpdateQuery("users", user, []string{"ConfirmedNumber"})
	fmt.Println("query :", query)

}

func GenerateInsertOrUpdateQuery(tableName string, data interface{}, updateColumns []string) string {
	valueType := reflect.TypeOf(data)
	// value := reflect.ValueOf(data)

	if valueType.Kind() != reflect.Struct {
		return ""
	}

	var columns []string

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fieldName := field.Name
		columns = append(columns, fieldName)
	}

	columnsStr := strings.Join(columns, ", ")

	// 업데이트할 열과 값을 정확하게 지정합니다.
	var updateValues []string
	for _, col := range updateColumns {
		updateValues = append(updateValues, fmt.Sprintf("%s = '%s'", col, getColumnValue(data, col)))
	}
	updateValuesStr := strings.Join(updateValues, ", ")

	// 단일 SQL 문으로 INSERT 또는 UPDATE 쿼리 생성
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s", tableName, columnsStr, getValuesString(data), updateValuesStr)

	return query
}

// getColumnValue 함수는 struct의 특정 열의 값을 반환합니다.
func getColumnValue(data interface{}, columnName string) string {
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Struct {
		field := value.FieldByName(columnName)
		if field.IsValid() {
			return fmt.Sprintf("%v", field.Interface())
		}
	}
	return ""
}

// getValuesString 함수는 struct의 값들을 쉼표로 구분된 문자열로 반환합니다.
func getValuesString(data interface{}) string {
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Struct {
		var values []string
		for i := 0; i < value.NumField(); i++ {
			values = append(values, getColumnValue(data, value.Type().Field(i).Name))
		}
		return strings.Join(values, ", ")
	}
	return ""
}
