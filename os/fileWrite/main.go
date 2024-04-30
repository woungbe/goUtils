package main

import (
	"encoding/json"
	"fmt"
	"os"

	"woungbe.utils/os/defined"
)

func main() {

	users := defined.Users{}
	err := defined.FileRead("../defined/users.json", &users)
	if err != nil {
		fmt.Println("어라 파일이 없네요 ?")
		return
	}

	b := true
	for _, v := range users.Users {
		if v.UserIDX == 1029 {
			b = false
		}
	}

	if b {
		user := defined.User{
			UserIDX: 1029, UserID: "selab.park@gmail.com", Passwd: "1234",
		}
		users.Users = append(users.Users, user)
	}

	err1 := WriteFile("../defined/users.json", &users)
	if err1 != nil {
		fmt.Println("err1 : ", err1)
		return
	}

}

func WriteFile(filename string, aa interface{}) error {
	// JSON 형식으로 변환
	jsonData, err := json.MarshalIndent(aa, "", "	")
	if err != nil {
		// panic(err)
		return err
	}

	// 파일에 기록
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
