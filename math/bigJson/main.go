package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Sam struct {
	Column json.Number `json:"column"`
	Fee    int64
}

func main() {

	src := strings.NewReader(`{
		"column": "972804139105955"
	}`)

	var inf Sam
	err := json.NewDecoder(src).Decode(&inf)
	if err != nil {
		return
	}

	fmt.Println("inf : ", inf.Column)

}
