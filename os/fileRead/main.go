package main

import (
	"fmt"

	"woungbe.utils/os/defined"
)

func main() {
	// 파일 읽기
	file := "./defined/users.json"
	users := defined.Users{}
	defined.FileRead(file, &users)
	fmt.Printf("%+v\n", users)

	configFile := "./defined/config.json"
	configs := defined.MMConfigCtrl{}
	defined.FileRead(configFile, &configs)
	fmt.Printf("%+v\n", &configs)
}
