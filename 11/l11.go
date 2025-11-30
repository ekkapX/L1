package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы, присутствующие и в первом, и во втором.

Пример:
A = {1,2,3}
B = {2,3,4}
Пересечение = {2,3}
*/

func intersection(a, b []int) []int {
	m := make(map[int]bool)
	res := make([]int, 0, len(a))
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if m[v] {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Println(intersection([]int{1, 2, 3}, []int{2, 3, 4})) // {2, 3}
	fmt.Println(intersection([]int{1, 2, 3}, []int{4, 5, 6})) // {}
}
