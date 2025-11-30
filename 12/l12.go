package main

import "fmt"

/*
Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.
*/

func set(nums []string) []string {
	seen := make(map[string]struct{}, len(nums))
	res := make([]string, 0, len(nums))
	for _, elem := range nums {
		if _, ok := seen[elem]; !ok {
			seen[elem] = struct{}{}
			res = append(res, elem)
		}
	}
	return res
}

func main() {
	fmt.Println(set([]string{"cat", "cat", "dog", "cat", "tree"}))
	fmt.Println(set([]string{"cat", "dog", "tree"}))
	fmt.Println(set([]string{"cat", "cat", "cat", "cat", "cat"}))

}
