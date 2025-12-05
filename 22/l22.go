package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

Комментарий: в Go тип int справится с такими числами, но обратите внимание на возможное переполнение для ещё больших значений. Для очень больших чисел можно использовать math/big.
*/

func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func subtract(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func divide(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {
	a := big.NewInt(2 << 20)
	b := big.NewInt(2 << 20)
	fmt.Println(add(a, b))
	fmt.Println(subtract(a, b))
	fmt.Println(multiply(a, b))
	fmt.Println(divide(a, b))
}
