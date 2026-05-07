package main

import (
	"fmt"
	"sync"
)

func fibonacci(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return fibonacci(num-2) + fibonacci(num-1)
}

func worker(job, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	Outer:
		for {
			select {
			case v, ok := <-job:
				if ok {
					result <- fibonacci(v)
				} else {
					break Outer
				}
			}
		}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	wg := sync.WaitGroup{}

	for range 4 {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	go func() {
		for v := range 10 {
			jobs <- v + 1
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for val := range results {
		fmt.Println("Values are ", val)
	}
}
