package main

import "fmt"

func main() {

	whereDate := "`Date`"
	StartDate := "2024-01-01"
	endDate := "2024-03-01"
	sql := fmt.Sprintf(`SELECT Symbol, sum(TraderCnt) as TraderCnt, count(UserIDX) as TraderUsers, sum(TraderPrice) as TraderPrice, 
	sum(QtyCnt) as QtyCnt, sum(CommissionAsset) as CommissionAsset, sum(CommissionUSDT) as CommissionUSDT
	FROM Batch_Spot_Trade where %s between '%s' and '%s'`, whereDate, StartDate, endDate)

	fmt.Println(sql)

}

// 회사에 앉아 있는다고 돈주는거면 ... 말이 안되지 . 


하루에 최소 100만원씩 하려면. 

2000 넣으면 일주일 평균 700



