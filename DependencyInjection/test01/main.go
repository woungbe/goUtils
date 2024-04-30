package main

import "fmt"

type Logger interface {
	Log(message string)
}
type MyService struct {
	logger Logger
}

func NewMyService(logger Logger) *MyService {
	return &MyService{
		logger: logger,
	}
}

func (s *MyService) DoSomething() {
	s.logger.Log("Doing something...")
}

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

type FmtLogger struct{}

func (l FmtLogger) Log(message string) {
	fmt.Println("from : ", message)
}

func CheckBind(logtype string) {
	if logtype == "1" {
		logger := ConsoleLogger{}
		service := NewMyService(logger)
		service.DoSomething()
	} else {
		flogger := FmtLogger{}
		service2 := NewMyService(flogger)
		service2.DoSomething()
	}
}

func main() {
	CheckBind("2")
}
