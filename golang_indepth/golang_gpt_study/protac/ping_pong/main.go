package main

import (
	"fmt"
	"sync"
)

func pingWorker(ch chan string, wg *sync.WaitGroup, count int) {
	defer wg.Done()
	for {
		val, ok := <- ch
		if !ok {
			fmt.Println("Channel closed")
			return
		}
		if count >= 5 {
			close(ch)
			return
		}
		fmt.Println(val)
		ch <- "Ping"
		count ++
	}
}

func pongWorker(ch chan string, wg *sync.WaitGroup, count int) {
	defer wg.Done()
	for {
		val, ok := <- ch
		if !ok {
			fmt.Println("Channel closed")
			return
		}
		if count >= 5 {
			close(ch)
			return
		}
		fmt.Println(val)
		ch <- "Pong"
		count ++
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan string)
	count := 0

	wg.Add(2)
	go pingWorker(ch, wg, count)
	go pongWorker(ch, wg, count)

	ch <- "Ping"

	wg.Wait()
}

















