package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.

Подсказка: используйте time.After или таймер для ограничения времени работы.
*/

func main() {
	fmt.Print("Введите время в секундах: ")
	var n int
	fmt.Scan(&n)
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	timeout := time.After(time.Second * time.Duration(n))

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-timeout:
			fmt.Println("Время вышло, завершение работы")
			return
		}
	}
}
