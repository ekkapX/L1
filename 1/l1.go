package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
	ActionName string
}

func (h Human) Do() {
	fmt.Println(h.Name + " doing something")
}

func (a Action) Info() {
	fmt.Printf("His hame %s, he's %d years old and his action actually doing %s", a.Name, a.Age, a.ActionName)
}

func main() {
	action := Action{Human: Human{Name: "Dmitro", Age: 23}, ActionName: "nothing"}
	action.Do()
	action.Info()
}
