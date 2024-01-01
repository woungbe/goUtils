package main

import "fmt"

type Pack interface {
	Hello() string
	World() string
}

type AA struct {
}

func (ty *AA) Hello() string {
	return "hello"
}

func (ty *AA) World() string {
	return "world"
}

type Ctrl struct {
	C map[string]Pack
}

func (ty *Ctrl) Init() {
	ty.C = make(map[string]Pack)
}

func main() {
	tt := new(Ctrl)
	tt.Init()

	aa := new(AA)
	tt.C["Pack"] = aa
	fmt.Println(tt.C["Pack"].Hello(), tt.C["Pack"].World())

}
