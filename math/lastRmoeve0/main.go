package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "0.01000000"
	trimmedStr := strings.TrimRight(str, "0")

	fmt.Println(trimmedStr)
}
