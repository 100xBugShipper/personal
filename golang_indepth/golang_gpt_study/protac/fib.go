package protac

import (
	"fmt"
	"sync"
)

func fib(i int) int {
	if i == 1 {
		return 1
	}else {
		return fib(i-2) + fib(i-1)
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 100)

	i := 0

	for {
		ch <- i
		i ++

		if i > 100 {
			break
		}
	}

	num := 0

	for range ch {
		num = <- ch
		fmt.Println(num)
	}
	for range 4 {
		wg.Add(1)
		go func(){
			fmt.Println(fib(num))
		}()
		wg.Done()
	}
	wg.Wait()

}
