package main

import (
	"fmt"
	"sort"
	"strconv"
)

type RESTV2_Inno_CoinListByUserOut struct {
	UserIDX           string `json:"userID" db:"UserIDX"`
	UID               string `json:"UID" db:"UID"`
	Symbol            string `json:"symbol" db:"Symbol"`
	Balance           string `json:"balance" db:"Balance"`
	LockedBal         string `json:"lockedBal" db:"LockedBal"`
	WithdrawLockedBal string `json:"withdrawLockedBal" db:"WithdrawLockedBal"`
	FreeBal           string `json:"freeBal" db:"FreeBal"`
}

// SortCoinsByBalanceDesc 함수는 send 슬라이스를 Balance 기준으로 내림차순 정렬합니다.
func SortCoinsByBalanceDesc(send []RESTV2_Inno_CoinListByUserOut) {
	sort.Slice(send, func(i, j int) bool {
		iBalance, errI := strconv.ParseFloat(send[i].Balance, 64)
		jBalance, errJ := strconv.ParseFloat(send[j].Balance, 64)
		if errI != nil && errJ != nil {
			return false // 두 값 모두 변환에 실패한 경우 순서 변경 없음
		}
		if errI != nil {
			return false // i번째 값 변환 실패, j가 더 크다고 간주
		}
		if errJ != nil {
			return true // j번째 값 변환 실패, i가 더 크다고 간주
		}
		return iBalance > jBalance // 실제 숫자 기준으로 내림차순 정렬
	})
}

func main() {
	// 테스트 데이터
	send := []RESTV2_Inno_CoinListByUserOut{
		{Balance: "100.5"},
		{Balance: "300.25"},
		{Balance: "50.75"},
		{Balance: "200.0"},
		{Balance: "500.12"},
	}

	// Balance 기준으로 내림차순 정렬
	SortCoinsByBalanceDesc(send)

	// 정렬된 결과 출력
	for _, coin := range send {
		fmt.Printf("Balance: %s\n", coin.Balance)
	}
}
