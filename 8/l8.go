package main

import "fmt"

/*

Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.

Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).

Подсказка: используйте битовые операции (|, &^).
*/

func setBit(value int64, i uint, bit int) int64 {
	if bit == 1 {
		return value | (1 << i)
	}
	return value &^ (1 << i)
}

func main() {
	fmt.Println(setBit(5, 1, 0))
	fmt.Println(setBit(5, 0, 0))
}
