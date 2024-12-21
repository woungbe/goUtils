package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey = "b465887887dac8"
const apiURL = "https://ipinfo.io/batch?token=" + apiKey

type IPInfo struct {
	Country string `json:"country"`
}

func getIPInfo(ipAddresses []string) (map[string]IPInfo, error) {
	// Create JSON payload
	payload, err := json.Marshal(ipAddresses)
	if err != nil {
		return nil, err
	}

	// Make the HTTP request
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var result map[string]IPInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func batchIPInfo(ipAddresses []string, batchSize int) (map[string]IPInfo, error) {
	allResults := make(map[string]IPInfo)

	for i := 0; i < len(ipAddresses); i += batchSize {
		end := i + batchSize
		if end > len(ipAddresses) {
			end = len(ipAddresses)
		}
		batch := ipAddresses[i:end]
		result, err := getIPInfo(batch)
		if err != nil {
			return nil, err
		}
		for ip, info := range result {
			allResults[ip] = info
		}
	}

	return allResults, nil
}

func main() {
	// Example IP addresses
	ipAddresses := []string{"221.139.79.21", "221.139.79.52", "218.234.149.10", "221.150.109.88"}

	type DataList struct {
		IP  []string
		Cnt int
	}

	// Get IP information
	list := make(map[string]DataList)

	// Batch size
	batchSize := 1000

	// Get IP information in batches
	ipInfo, err := batchIPInfo(ipAddresses, batchSize)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	for ip, info := range ipInfo {
		fmt.Printf("IP: %s, Country: %s\n", ip, info.Country)

		if val, ok := list[info.Country]; ok {
			val.Cnt = val.Cnt + 1
			val.IP = append(val.IP, ip)
			list[info.Country] = val
		} else {
			val.IP = append(val.IP, ip)
			val.Cnt = 1
			list[info.Country] = val
		}
	}

	fmt.Printf("%+v\n", list)

	// Print the result
	/*
		for ip, info := range ipInfo {
			fmt.Printf("IP: %s, Country: %s\n", ip, info.Country)
		}
	*/

}

/*
한달 동안 거래한 총 유저 482 명
그중에서 KYC 한 유저 448 명
그 안에서 확인한 아이피 268명


국가별 접속 유저 입니다.

AU: Australia (호주)     2
JP: Japan (일본)         9
KR: South Korea (대한민국)  264
MY: Malaysia (말레이시아)   1
NZ: New Zealand (뉴질랜드)  1
*/

// 아이피 접속
