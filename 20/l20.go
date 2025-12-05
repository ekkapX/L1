package main

import (
	"fmt"
)

/*
Разработать программу, которая переворачивает порядок слов в строке.

Пример: входная строка:

«snow dog sun», выход: «sun dog snow».

Считайте, что слова разделяются одиночным пробелом. Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».
*/
func reverse(s []rune, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	res := []rune(s)
	reverse(res, 0, len(res)-1)

	l := 0
	for r := 0; r <= len(res); r++ {
		if r == len(res) || res[r] == ' ' {
			reverse(res, l, r-1)
			l = r + 1
		}
	}
	return string(res)
}

func main() {
	fmt.Println(reverseWords("snow dog sun"))
	fmt.Println(reverseWords("a good example"))
}
