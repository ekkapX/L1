package main

import (
	"fmt"
)

/*
Поменять местами два числа без использования временной переменной.

Подсказка: примените сложение/вычитание или XOR-обмен.
*/

func swap(x, y int) (int, int) {
	x ^= y
	y ^= x
	x ^= y
	return x, y
}

func main() {
	fmt.Println(swap(1, 2))
}
