package main

import "fmt"

type Users struct {
	Name string
}

type ChangeUser Users

func main() {

	user := Users{Name: "help"}

	SetUser(user)

	// fmt.Println(user)
	// Users{Name:}
}

func SetUser(agrs ChangeUser) {
	fmt.Println("agrs.Name : ", agrs.Name)
}
