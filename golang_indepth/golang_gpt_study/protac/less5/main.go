package main

import (
	"fmt"
	"sync"
)

func worker(odd, even, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	Outer:
	for {
		select {
		case val, ok := <-odd:
			if ok {
				results <- val
			}else {
				break Outer
			}
		case val, ok := <- even:
			if ok {
				results <- val
			}else {
				break Outer
			}
		}
	}
}

func main() {
	oddJobs := make(chan int)
	evenJobs := make(chan int)
	results := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go worker(oddJobs, evenJobs, results, &wg)

	go func() {
		for val := range 10{
			if val%2 == 0 {
				evenJobs <- val
			} else {
				oddJobs <- val
			}
		}
		close(oddJobs)
		close(evenJobs)
	}()

	go func() {
		wg.Wait()
		close(results)

	}()

	for val := range results {
		fmt.Println("Here are the values ", val)
	}

}
