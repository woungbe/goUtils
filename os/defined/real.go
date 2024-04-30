package defined

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	UserIDX int64
	UserID  string
	Passwd  string
}

type Users struct {
	Users []User
}

type MMConfig struct {
	UserIDX        int    `json:"UserIDX"`
	Symbol         string `json:"Symbol"`
	MMCountMin     int    `json:"MMCountMin"`
	MMCountMax     int    `json:"MMCountMax"`
	MMTimeMin      int    `json:"MMTimeMin"`
	MMTimeMax      int    `json:"MMTimeMax"`
	MMCostRangeMin int    `json:"MMCostRangeMin"`
	MMCostRangeMax int    `json:"MMCostRangeMax"`
	MMOPTMaxCount  int    `json:"MMOPTMaxCount"`
	MMOPTType      string `json:"MMOPTType"`
	MMType         string `json:"MMType"`
	MMStatus       int    `json:"MMStatus"`
	ABRangeTickMin int    `json:"ABRangeTickMin"`
	ABRangeTickMax int    `json:"ABRangeTickMax"`
	ABCountMin     int    `json:"ABCountMin"`
	ABCountMax     int    `json:"ABCountMax"`
	ABTimeMin      int    `json:"ABTimeMin"`
	ABTimeMax      int    `json:"ABTimeMax"`
	ABotStatus     int    `json:"ABotStatus"`
}

type MMConfigCtrl struct {
	Config []MMConfig `json:"Config"`
}

func FileRead(files string, a interface{}) error {
	b, err := os.ReadFile(files)
	if err != nil {
		return err
	}

	er := json.Unmarshal(b, a)
	if er != nil {
		fmt.Println(er)
		return er
	}
	return nil
}
