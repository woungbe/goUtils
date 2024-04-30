package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// DebounceFunction 타입은 어떤 함수라도 받을 수 있도록 제네릭으로 정의됩니다.
type DebounceFunction[T any] T

// Debounce 함수는 제네릭 함수 타입 F를 매개변수로 받습니다.
func Debounce[F any](fn F, duration time.Duration) F {
	var mutex sync.Mutex
	var timer *time.Timer
	var once sync.Once

	// 함수를 호출하기 위한 변수
	var callFn func(args []any)

	// 제네릭 함수를 호출하기 위한 함수
	callFn = func(args []any) {
		fnVal := reflect.ValueOf(fn)
		fnArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			fnArgs[i] = reflect.ValueOf(arg)
		}
		fnVal.Call(fnArgs)
	}

	// 결과 함수는 interface{} 슬라이스를 매개변수로 받습니다.
	var debouncedFn any = func(args ...any) {
		mutex.Lock()
		defer mutex.Unlock()

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(duration, func() {
			mutex.Lock()
			defer mutex.Unlock()

			once.Do(func() {
				callFn(args)
			})
		})
	}

	return reflect.MakeFunc(reflect.TypeOf(fn), debouncedFn.(func([]reflect.Value) []reflect.Value)).Interface().(F)
}

func main() {
	// 사용 예시: 매개변수가 있는 함수
	testFunc := func(msg string, name string) {
		fmt.Println("Function called with:", msg, name)
	}

	// Debounce 적용
	debouncedFunc := Debounce(testFunc, 2*time.Second)

	// 여러번 호출
	for i := 0; i < 5; i++ {
		debouncedFunc("Hello, world!", "경령")
		time.Sleep(500 * time.Millisecond)
	}

	// 출력 결과 확인을 위해 대기
	time.Sleep(3 * time.Second)
}
