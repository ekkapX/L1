package main

import (
	"fmt"
	"slices"
)

/*
Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.

Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.
*/

func delete(a []int, i int) []int {
	if i < 0 || i >= len(a) {
		return a
	}
	copy(a[i:], a[i+1:])
	return a[:len(a)-1]
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	a = slices.Delete(a, 2, 3) // [1 2 4 5]
	fmt.Println(a)
	a = delete(a, 2) // [1 2 5]
	fmt.Println(a)
}
