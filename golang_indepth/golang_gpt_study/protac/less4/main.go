package main

import (
	"fmt"
	"sync"
)

func calcFibonacci(num int) int {
	if num <= 1 {
		return 1
	}
	return calcFibonacci(num - 2) + calcFibonacci(num - 1)
}

func worker(k int, job, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("I am worker %d\n", k)
	for val := range job {
		result <- calcFibonacci(val)
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	wg := sync.WaitGroup{}


	for k := range 4 {
		wg.Add(1)
		go worker(k+1, jobs, results, &wg)
	}

	for i := range 10 {
		jobs <- i + 1
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for val := range results {
		fmt.Printf("Here are the results -> %d\n", val)
	}

}
