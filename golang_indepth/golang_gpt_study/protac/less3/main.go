package main

import (
	"fmt"
	"sync"
)

func worker1(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for range 5 {
		fmt.Println("Inside the worker 1")
		val := <-ch
		fmt.Println(val)
		ch <- "Ping"
	}
}

func worker2(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for range 5 {
		fmt.Println("Inside the worker 1")
		val := <-ch
		fmt.Println(val)
		ch <- "Pong"
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(1)
	go worker1(&wg, ch)
	wg.Add(1)
	go worker2(&wg, ch)

	ch <- "Start"
	go func() {
		wg.Wait()
		close(ch)
	}()
}
