package main

import "fmt"

// OnewayBot 구조체 정의
type OnewayBot struct {
	ID             int
	BinanceAccount *BinanceAccount
}

// RemoveOnewayBot implements IBot.
func (ba *OnewayBot) RemoveOnewayBot(int) {
	panic("unimplemented")
}

func (ba *OnewayBot) GetID() int {
	return ba.ID
}

type IBot interface {
	GetID() int
	RemoveOnewayBot(int)
}

// BinanceAccount 구조체 정의
type BinanceAccount struct {
	ID   int
	Bots []IBot
}

// OnewayBot 제거 함수
func (ba *BinanceAccount) RemoveOnewayBot(botID int) {
	newBots := []IBot{}
	for _, bot := range ba.Bots {
		if bot.GetID() != botID {
			newBots = append(newBots, bot)
		}
	}
	ba.Bots = newBots
}

func (ba *BinanceAccount) GetID() int {
	return ba.ID
}

func main() {
	// BinanceAccount 인스턴스 생성
	account := &BinanceAccount{}

	// OnewayBot 인스턴스 생성 및 BinanceAccount 설정
	bot1 := &OnewayBot{ID: 1, BinanceAccount: account}
	bot2 := &OnewayBot{ID: 2, BinanceAccount: account}
	bot3 := &OnewayBot{ID: 3, BinanceAccount: account}

	// BinanceAccount에 OnewayBot 추가
	account.Bots = append(account.Bots, bot1, bot2, bot3)

	// OnewayBot 제거 전 상태 출력
	fmt.Println("Before removing bot:")
	for _, bot := range account.Bots {
		fmt.Printf("OnewayBot ID: %d\n", bot.GetID())
	}

	// OnewayBot 제거
	account.RemoveOnewayBot(2)

	// OnewayBot 제거 후 상태 출력
	fmt.Println("After removing bot:")
	for _, bot := range account.Bots {
		fmt.Printf("OnewayBot ID: %d\n", bot.GetID())
	}
}
