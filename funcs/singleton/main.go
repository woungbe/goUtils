package main

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	instance *Singleton
)

// Singleton 구조체 예제
type Singleton struct {
	Value string
}

// GetInstance 함수는 Singleton 인스턴스를 반환하며, 단 한 번만 생성됨을 보장
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating Singleton instance")
		instance = &Singleton{Value: "My Singleton"}
	})
	return instance
}

func main() {
	// 여러 고루틴에서 GetInstance 함수를 호출하여도 Singleton 인스턴스는 한 번만 생성됨
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			singleton := GetInstance()
			fmt.Printf("Goroutine %d: Singleton instance value: %s\n", id, singleton.Value)
		}(i)
	}

	wg.Wait()
}
