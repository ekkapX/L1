package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).

Вывод: true, если все символы уникальны, false, если есть повторения. Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.

Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.

Подумайте, какой структурой данных удобно воспользоваться для проверки условия.
*/

func allCharsUnique(s string) bool {
	seen := make(map[string]struct{}, len(s))
	for _, elem := range s {
		if _, ok := seen[strings.ToLower(string(elem))]; !ok {
			seen[strings.ToLower(string(elem))] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(allCharsUnique("abcd"))
	fmt.Println(allCharsUnique("abCdefAaf"))
	fmt.Println(allCharsUnique("aabcd"))
}
