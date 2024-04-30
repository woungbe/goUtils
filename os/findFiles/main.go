package main

import (
	"log"
	"os"
)

func main() {
	filename := "example.txt"

	// 파일 정보 가져오기
	_, err := os.Stat(filename)

	// 파일이 없는 경우 에러 처리
	if os.IsNotExist(err) {
		log.Printf("파일 %s가 존재하지 않습니다.", filename)
		return
	}

	log.Printf("파일 %s가 존재합니다.", filename)
}
