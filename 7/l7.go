package main

import (
	"fmt"
	"sync"
)

/*
Реализовать безопасную для конкуренции запись данных в структуру map.

Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).

Проверьте работу кода на гонки (util go run -race).
*/

type syncMap struct {
	m  map[int]int
	mu sync.RWMutex
}

func newSyncMap() *syncMap {
	return &syncMap{
		m: make(map[int]int),
	}
}

func (s *syncMap) Set(key, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func main() {
	sm := newSyncMap()
	wg := sync.WaitGroup{}

	wg.Add(10)
	for i := range 10 {
		go func(i int) {
			sm.Set(i, i)
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	sm.mu.Lock()
	fmt.Println(sm.m)
	sm.mu.Unlock()
}
