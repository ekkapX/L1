package main

import "fmt"

func typeOf(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int, chan string, chan bool:
		return "chan int"
	default:
		return "unknown"
	}

}
func main() {
	fmt.Println(typeOf(42))
	fmt.Println(typeOf("hello"))
	fmt.Println(typeOf(true))
	fmt.Println(typeOf(make(chan int)))
	fmt.Println(typeOf(1.2))
}
