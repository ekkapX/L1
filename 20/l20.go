package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	words := strings.Split(s, " ")
	reversedWords := make([]string, len(words))
	for i := len(words) - 1; i >= 0; i-- {
		reversedWords = append(reversedWords, words[i])
	}
	return strings.Join(reversedWords, " ")
}

func main() {
	fmt.Println(reverse("hello world"))
	fmt.Println(reverse("snow dog sun"))
}
