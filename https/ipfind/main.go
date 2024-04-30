package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("서버 시작!!")
	http.HandleFunc("/", LimitorForMM_HandleFunc(handler, false)) // 루트 경로 핸들러 설정
	http.ListenAndServe(":9910", nil)                             // 서버 시작

}

func handler(w http.ResponseWriter, r *http.Request) {
	// 쿼리 파라미터 추출
	query := r.URL.Query()
	// name := query.Get("name") // 'name' 파라미터 값 추출
	symbol := query.Get("symbol")
	interval := query.Get("interval")
	startTime := query.Get("startTime")
	endTime := query.Get("endTime")
	limit := query.Get("limit")

	// 응답으로 파라미터 값 반환
	fmt.Fprintln(w, "Hello, ", symbol, interval, startTime, endTime, limit)
}

func LimitorForMM_HandleFunc(next func(w http.ResponseWriter, r *http.Request), notCheck bool) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("==================", r.URL.Path)
		// cnf := svrconfig.GetConfigData()
		if notCheck {
			next(w, r)
			return
		}

		// strIP, _, err := net.SplitHostPort(r.RemoteAddr)
		// if err != nil {
		// 	strIP = r.RemoteAddr
		// }

		realip := GetRealIP(r)
		fmt.Println("1111111111111111111111")
		fmt.Println("realip : ", realip)

		b, _ := IsIPInRange(realip, "172.31.0.0/16")
		fmt.Println("b :", b)
		if !b {
			next(w, r)
			return
		}

		next(w, r)
	})
}

func GetRealIP(r *http.Request) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Crit Error [ respLimitor.GetRealIP ]")
		}
	}()

	var ip string
	if tcip := r.Header.Get("True-Client-IP"); tcip != "" {
		ip = tcip
	} else if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		ip = xrip
	} else if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		i := strings.Index(xff, ", ")
		if i == -1 {
			i = len(xff)
		}
		ip = xff[:i]
	} else {
		var err error
		ip, _, err = net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}
	}
	return canonicalizeIP(ip)
}

func canonicalizeIP(ip string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Crit Error [ respLimitor.canonicalizeIP ]")
		}
	}()
	isIPv6 := false
	// This is how net.ParseIP decides if an address is IPv6
	// https://cs.opensource.google/go/go/+/refs/tags/go1.17.7:src/net/ip.go;l=704
	for i := 0; !isIPv6 && i < len(ip); i++ {
		switch ip[i] {
		case '.':
			// IPv4
			return ip
		case ':':
			// IPv6
			isIPv6 = true
			break
		}
	}
	if !isIPv6 {
		// Not an IP address at all
		return ip
	}

	ipv6 := net.ParseIP(ip)
	if ipv6 == nil {
		return ip
	}

	return ipv6.Mask(net.CIDRMask(64, 128)).String()
}

// IsIPInRange 함수는 주어진 IP가 지정된 CIDR 범위 내에 있는지 확인합니다.
func IsIPInRange(ipStr, cidrStr string) (bool, error) {
	ip := net.ParseIP(ipStr)
	_, cidrNet, err := net.ParseCIDR(cidrStr)
	if err != nil {
		return false, err
	}
	return cidrNet.Contains(ip), nil
}
