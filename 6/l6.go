package main

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.

Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.

Продемонстрируйте каждый способ в отдельном фрагменте кода.
*/

func main() {

	// 1. Выход с помощью канала
	ch := make(chan struct{})
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("1. Goroutine closed")
				return
			default:
				fmt.Println("1. Goroutine working")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	close(ch)
	time.Sleep(time.Second * 1)

	// 2. Выход с помощью runtime.Goexit
	go func() {
		fmt.Println("2. Goroutine working")
		defer fmt.Println("2. Goroutine closed")
		runtime.Goexit()
	}()

	time.Sleep(time.Second * 2)

	// 3. Выход при помощи флага
	var flag atomic.Bool
	go func() {
		for {
			if flag.Load() {
				fmt.Println("3. Goroutine closed")
				return
			}
			fmt.Println("3. Goroutine working")
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 2)
	flag.Store(true)
	time.Sleep(time.Second * 1)

	// 4. Выход с помощью контекста с отменой
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("4. Goroutine closed")
				return
			default:
				fmt.Println("4. Goroutine working")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second * 1)

	// 5. Выход с помощью контекста с тайм-аутом
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Duration(time.Second*2))
	defer cancel2()
	go func() {
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("5. Goroutine closed")
				return
			default:
				fmt.Println("5. Goroutine working")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(time.Second * 3)
}
