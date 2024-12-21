package main

import (
	"fmt"
	"time"
)

func main() {

	doen := make(chan (bool))
	go func() {
		i := 0
		for {
			i++
			fmt.Println(i, ":출력해주세요 ")
			time.Sleep(Init(5))
		}
	}()

	<-doen
}

func Init(aftime int64) time.Duration {
	AfterKYCTime := time.Duration(aftime) * time.Second
	return AfterKYCTime
}
