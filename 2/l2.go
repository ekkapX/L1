package main

import (
	"fmt"
	"sync"
)

func main() {
	var numbers = [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for i := range numbers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			numbers[i] *= numbers[i]
			fmt.Println(numbers[i])
		}(i)
	}
	wg.Wait()
}
