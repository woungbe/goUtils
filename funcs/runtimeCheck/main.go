package main

import (
	"fmt"
	"os"
	"strings"
)

/*
go run main.go 로 실행했는지

build 로 실행하는지 확인하는 방법


*/

func main() {
	executable, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("executable : ", executable)

	// Check for the DEBUG_MODE environment variable
	debugMode := os.Getenv("DEBUG_MODE")
	if strings.Contains(executable, "go-build") {
		fmt.Println("Running with 'go run'")
	} else if debugMode == "true" {
		fmt.Println("Running in debug mode")
	} else {
		fmt.Println("Running from a built executable")
	}

	fmt.Println("Executable path:", executable)

}
