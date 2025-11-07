package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).

Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.

Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done()
		close(jobs)
	}()

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
