package main

import (
	"fmt"
	"strconv"
)

func formatPrice(precision int, price string) string {
	// 문자열을 실수로 변환
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return "" // 변환에 실패한 경우 빈 문자열 반환
	}

	// 실수를 포맷하여 출력
	return fmt.Sprintf("%.*f", precision, priceFloat)
}

func main() {
	BaseAssetPrecision := 6
	PriceSizeInfo_MinPrice := "0.001"  // string 타입으로 변경
	PriceSizeInfo_MaxPrice := "900000" // string 타입으로 변경

	// 최소가격 출력
	minPriceFormatted := formatPrice(BaseAssetPrecision, PriceSizeInfo_MinPrice)
	fmt.Println(minPriceFormatted)

	// 최대가격 출력
	maxPriceFormatted := formatPrice(BaseAssetPrecision, PriceSizeInfo_MaxPrice)
	fmt.Println(maxPriceFormatted)
}
