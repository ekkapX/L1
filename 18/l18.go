package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин). По завершению программы структура должна выводить итоговое значение счётчика.

Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.
*/

type syncCounter struct {
	count int
	mu    sync.Mutex
}

func (c *syncCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func main() {
	var wg sync.WaitGroup
	counter := syncCounter{}

	wg.Add(3)
	for range 3 {
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.Increment()
			}
		}()
	}
	wg.Wait()

	fmt.Println(counter.count)
}
