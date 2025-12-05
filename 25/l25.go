package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep, которая приостанавливает выполнение текущей горутины.

Важно: в отличие от настоящей time.Sleep, ваша функция должна именно блокировать выполнение (например, через таймер или цикл), а не просто вызывать time.Sleep :) — это упражнение.
*/

func sleep(t time.Duration) {
	ch := make(chan struct{})
	go func() {
		timer := time.NewTimer(t)
		<-timer.C
		close(ch)
	}()
	<-ch
}

func main() {
	now := time.Now()
	sleep(5 * time.Second)
	fmt.Println(time.Since(now))
}
