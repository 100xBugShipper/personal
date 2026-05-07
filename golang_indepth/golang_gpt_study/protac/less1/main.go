package main

import (
	"fmt"
	"sync"
)


func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		ch <- i * i
	}
}

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go worker(ch, wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("Values :- %d\n", i)
	}
}
