package main

import (
	"fmt"
	"net"
)

func FindServerIP() []string {
	var send []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return send
	}
	for _, addr := range addrs {
		// 네트워크 주소를 검사하여 IP 네트워크 주소인지 확인합니다.
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			// IPv4 주소인지 확인합니다.
			if ipNet.IP.To4() != nil {
				// fmt.Println("IPv4: ", ipNet.IP.String())
				send = append(send, ipNet.IP.String())
			}
		}
	}
	return send
}

func main() {
	IPs := FindServerIP()
	fmt.Println(IPs)
}
