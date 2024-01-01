package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

var (
	kernel32DLL   = syscall.NewLazyDLL("kernel32.dll")
	setSystemTime = kernel32DLL.NewProc("SetSystemTime")
)

func main() {
	service()

	done := make(chan bool)
	r := bufio.NewReader(os.Stdin)
	go func() {
		for {
			line, err := r.ReadString('\n')
			if err != nil && err.Error() != "unexpected newline" {
				// fmt.Println(err.Error())
				fmt.Println()
				//	return
				line = ""
			}

			line = strings.TrimSpace(line)

		}
	}()
	<-done
}

func service() {
	// NTP 서버 주소와 포트 번호 설정
	ntpServer := "time.google.com:123"

	// UDP 연결을 생성합니다.
	conn, err := net.Dial("udp", ntpServer)
	if err != nil {
		fmt.Println("NTP 서버에 연결할 수 없습니다:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// NTP 패킷 구조를 정의합니다.
	// 첫 번째 바이트는 NTP 버전 및 모드를 나타냅니다.
	// 두 번째 바이트는 서버와 클라이언트 간의 시간 동기화에 사용됩니다.
	// 나머지 부분은 무시합니다.
	ntpPacket := make([]byte, 48)
	ntpPacket[0] = 0x1B

	// NTP 패킷을 서버로 전송합니다.
	_, err = conn.Write(ntpPacket)
	if err != nil {
		fmt.Println("NTP 패킷을 전송할 수 없습니다:", err)
		os.Exit(1)
	}

	// 서버로부터 시간 데이터를 읽습니다.
	_, err = conn.Read(ntpPacket)
	if err != nil {
		fmt.Println("시간 데이터를 읽을 수 없습니다:", err)
		os.Exit(1)
	}

	// NTP 패킷에서 시간 데이터 추출
	// 70년 1월 1일부터 현재까지의 초 단위 시간을 계산합니다.
	timestamp := uint64(ntpPacket[43])<<24 | uint64(ntpPacket[42])<<16 | uint64(ntpPacket[41])<<8 | uint64(ntpPacket[40])
	unixTime := int64(timestamp - 2208988800)

	// 시간을 형식화하여 출력
	currentTime := time.Unix(unixTime, 0)
	fmt.Println("서버 시간:", currentTime)

	// 로컬 시스템 시간을 업데이트
	err = setTime(currentTime)
	if err != nil {
		fmt.Println("로컬 시스템 시간을 업데이트할 수 없습니다:", err)
		os.Exit(1)
	}

	fmt.Println("로컬 시스템 시간이 성공적으로 동기화되었습니다.")
}

// 로컬 시스템 시간을 설정하는 함수
func setTime(newTime time.Time) error {
	// SYSTEMTIME 구조체를 생성하고 필요한 시간 정보를 설정합니다.
	var systemTime syscall.Systemtime
	systemTime.Year = uint16(newTime.Year())
	systemTime.Month = uint16(newTime.Month())
	systemTime.Day = uint16(newTime.Day())
	systemTime.Hour = uint16(newTime.Hour())
	systemTime.Minute = uint16(newTime.Minute())
	systemTime.Second = uint16(newTime.Second())

	// Windows API를 호출하여 시스템 시간을 설정합니다.
	r, _, err := setSystemTime.Call(uintptr(unsafe.Pointer(&systemTime)))
	if r == 0 {
		return err
	}

	return nil
}
