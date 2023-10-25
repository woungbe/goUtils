package main

import (
	"fmt"
	"math"
	"strconv"
)

// 최소금액 / 현재가 = 최소 수량
// 현재가 * 최소수량 = 최소 금액

// 둘다 만족해야됨
// (최소수량 * 최소금액) * 최소수량 = 최소금액

func main() {

	// 1e-05 => 0.00001 로 변경해주는것
	// float64 => string  으로 변경됨 !
	coinSize := 0
	// tmp := MinCoinSize("26524", "5", "0.00001", coinSize)

	tmp2 := MinCoinSize("0.5077", "5", "0.0001", coinSize)

	fmt.Println("tmp : ", tmp2)
	formatString := fmt.Sprintf("%%.%df\n", coinSize)
	result := fmt.Sprintf(formatString, tmp2)
	fmt.Println(result)

}

// 현재가, 최소금액
func MinCoinSize(curPrice, minMoney, minSize string, coinStep int) float64 {

	mCurPrice, err := strconv.ParseFloat(curPrice, 64)
	if err != nil {
		fmt.Println(err)
	}

	mMinMoney, err := strconv.ParseFloat(minMoney, 64)
	if err != nil {
		fmt.Println(err)
	}

	mMinSize, err := strconv.ParseFloat(minSize, 64)
	if err != nil {
		fmt.Println(err)
	}

	// mCoinStep, err := strconv.ParseFloat(coinStep, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	if mMinMoney/mCurPrice < mMinSize {
		return mMinSize
	} else {
		x := mMinMoney / mCurPrice
		cnt := math.Pow(10, float64(coinStep))
		return math.Ceil(x*cnt) / cnt
	}
}

func Cell(num float64, decimalCnt string) float64 {
	tmp, er := strconv.Atoi(decimalCnt)
	if er != nil {
		fmt.Println("roundToTwoDecimalPlaces :", er)
	}

	cnt := math.Pow(10, float64(tmp))
	return math.Floor(num*cnt) / cnt
}

func roundToTwoDecimalPlaces(num float64, decimalCnt string) float64 {
	tmp, er := strconv.Atoi(decimalCnt)
	if er != nil {
		fmt.Println("roundToTwoDecimalPlaces :", er)
	}

	cnt := math.Pow(10, float64(tmp))
	return math.Floor(num*cnt) / cnt
}
