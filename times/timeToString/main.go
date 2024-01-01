package main

import (
	"fmt"
	"time"
)

func main() {
	var unix1 int64 = 1700697599
	var unix2 int64 = 1700697599999
	var unix3 int64 = 1700697599999999

	time1, err1 := unixTimestampToTime(unix1)
	if err1 != nil {
		fmt.Println("err1 : ", err1)
	}
	time2, err2 := unixTimestampToTime(unix2)
	if err2 != nil {
		fmt.Println("err2 : ", err2)
	}
	time3, err3 := unixTimestampToTime(unix3)
	if err3 != nil {
		fmt.Println("err3 : ", err3)
	}

	fmt.Println(TimeTostring(time1))
	fmt.Println(TimeTostring(time2))
	fmt.Println(TimeTostring(time3))
}

// unixtime 을,  date string으로 만드는 것
func UnixToDate(unixTime int64) (string, error) {
	times, err := unixTimestampToTime(unixTime)
	if err != nil {
		return "", err
	}

	send := TimeTostring(times)
	return send, nil
}

func TimeTostring(t time.Time) string {
	formatted := t.Format("2006-01-02")
	return formatted
}

func unixTimestampToTime(unixTime int64) (time.Time, error) {
	// UNIX 시간 스탬프의 자릿수에 따라 처리
	switch {
	case unixTime >= 1e18: // 19 자리
		return time.Unix(unixTime/1e9, unixTime%1e9), nil
	case unixTime >= 1e15: // 16 자리
		return time.Unix(unixTime/1e6, unixTime%1e6*1e3), nil
	case unixTime >= 1e12: // 16 자리
		return time.Unix(unixTime/1e3, unixTime%1e3*1e3), nil
	case unixTime >= 1e9: // 10 자리
		return time.Unix(unixTime, 0), nil
	default:
		return time.Time{}, fmt.Errorf("Invalid UNIX timestamp")
	}
}
