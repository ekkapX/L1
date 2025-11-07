package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать постоянную запись данных в канал (в главной горутине).

Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.
*/

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("Worker", id, "processed job", j)
	}
}

func main() {
	numWorkers := flag.Int("workers", 5, "number of workers")
	flag.Parse()
	if *numWorkers < 1 {
		fmt.Println("Number of workers must be greater than 0")
		return
	}
	jobs := make(chan int)
	var wg sync.WaitGroup
	wg.Add(*numWorkers)

	for w := 1; w <= *numWorkers; w++ {
		go worker(w, jobs, &wg)
	}

	go func() {
		for j := 1; ; j++ {
			jobs <- j
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
