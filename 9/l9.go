package main

import (
	"fmt"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2. После этого данные из второго канала должны выводиться в stdout. То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка. Убедитесь, что чтение из второго канала корректно завершается.
*/

func intStream(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, num := range nums {
			ch <- num
		}
	}()
	return ch
}

func squareStream(in <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for num := range in {
			ch <- num * num
		}
	}()
	return ch
}

func main() {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8}
	stream := intStream(source)

	for num := range squareStream(stream) {
		fmt.Println(num)
	}

}
