package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println("Hello, World!")
	// var a uint64 = 10000000000000000000
	max, _ := strconv.ParseFloat("1000000.00000000", 64)
	min, _ := strconv.ParseFloat("0.01000000", 64)

	var float float64 = 0.0010000000

	if float > max {
		fmt.Println("float > max", float, max)
	} 	

	if float < min {
		fmt.Println("float < min", float, min)
	}
}
