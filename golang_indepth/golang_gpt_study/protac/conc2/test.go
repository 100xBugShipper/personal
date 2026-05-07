package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

func worker(wg *sync.WaitGroup, i int, ch chan int) {
	defer wg.Done()
	ch <- i * i
	fmt.Printf("I am goroutine:- %s\n", uuid.New())
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(11)))
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	for i := range 10 {
		wg.Add(1)
		go worker(&wg, i, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Printf("Square of the number %d \n", val)
	}
}


// package main
//
// import (
// 	"context"
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// 	"github.com/google/uuid"
// )
//
// func worker(ctx context.Context, wg *sync.WaitGroup, i int, ch chan<- int) {
// 	defer wg.Done()
//
// 	select {
// 	case ch <- i * i:
// 		fmt.Printf("I am goroutine:- %s\n", uuid.New())
// 		time.Sleep(time.Millisecond * time.Duration(rand.Intn(11)))
// 	case <-ctx.Done():
// 		fmt.Printf("Worker %d cancelled\n", i)
// 		return
// 	}
// }
//
// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
//
// 	wg := sync.WaitGroup{}
// 	ch := make(chan int)
//
// 	// Start workers
// 	for i := range 10 {
// 		wg.Add(1)
// 		go worker(ctx, &wg, i, ch)
// 	}
//
// 	// Close channel when all workers are done
// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()
//
// 	// Read from channel with context cancellation support
// 	for {
// 		select {
// 		case val, ok := <-ch:
// 			if !ok {
// 				fmt.Println("Channel closed, all workers done")
// 				return
// 			}
// 			fmt.Printf("Square of numbers :- %d \n", val)
// 		case <-ctx.Done():
// 			fmt.Println("Context cancelled, shutting down")
// 			return
// 		}
// 	}
// }
