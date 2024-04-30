package main

import (
	"fmt"
	"strings"
)

// EscapeSQLString은 SQL 쿼리의 문자열 값에서 따옴표를 이스케이프합니다.
func EscapeSQLString(s string) string {
	return strings.ReplaceAll(s, "'", "")
}

func main() {
	// 예시 데이터
	resultID := "exampleResultID"
	withdrawDataInfo := struct {
		UserIDX      int
		Coin         string
		Network      string
		Address      string
		AddressTag   string
		Amount       string
		TxID         string
		ErrorMessage string
	}{
		UserIDX:      1,
		Coin:         "BTC",
		Network:      "Bitcoin",
		Address:      "1BitcoinAddress",
		AddressTag:   "Tag",
		Amount:       "0.001",
		TxID:         "Tx1234",
		ErrorMessage: "No woungbe's Error",
	}

	// SQL 쿼리 문자열 조립
	sqlL := fmt.Sprintf(" '%s', %d, '%s', '%s', '%s', '%s', '%s', 102, '%s', '%s' ",
		EscapeSQLString(resultID),
		withdrawDataInfo.UserIDX,
		EscapeSQLString(withdrawDataInfo.Coin),
		EscapeSQLString(withdrawDataInfo.Network),
		EscapeSQLString(withdrawDataInfo.Address),
		EscapeSQLString(withdrawDataInfo.AddressTag),
		EscapeSQLString(withdrawDataInfo.Amount),
		EscapeSQLString(withdrawDataInfo.TxID),
		EscapeSQLString(withdrawDataInfo.ErrorMessage))

	// sqlF는 여기서 정의되지 않았으나, 쿼리의 앞부분을 포함하고 있다고 가정합니다.
	// 여기에서 sql 변수를 DBQueryExec 함수에 전달하여 실행합니다.
	fmt.Println(sqlL)
}
