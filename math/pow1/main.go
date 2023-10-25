package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	val := roundToTwoDecimalPlaces(25424.2453, "2")
	fmt.Println("val : ", val)

}

func roundToTwoDecimalPlaces(num float64, decimalCnt string) float64 {
	tmp, er := strconv.Atoi(decimalCnt)
	if er != nil {
		fmt.Println("roundToTwoDecimalPlaces :", er)
	}

	cnt := math.Pow(10, float64(tmp))
	return math.Floor(num*cnt) / cnt
}
