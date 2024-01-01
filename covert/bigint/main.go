package main

import (
	"fmt"
	"math/big"
)

func main() {

	var i *big.Int
	i = big.NewInt(1000000000000000000)
	kk := fmt.Sprintf("%d", i)
	fmt.Println(kk)

}
