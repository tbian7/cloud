package main

import (
	"fmt"
	"sync"
)

func Split(src <-chan int, n int) []<-chan int {
	dests := make([]<-chan int, 0)

	for i := 0; i < n; i++ {
		dest := make(chan int)
		dests = append(dests, dest)

		go func() {
			defer close(dest)
			for n := range src {
				dest <- n
			}
		}()
	}

	return dests
}

func main() {
	src := make(chan int)

	go func() {
		defer close(src)

		for i := 0; i < 1000; i++ {
			src <- i
		}
	}()

	dests := Split(src, 10)

	var wg sync.WaitGroup
	wg.Add(len(dests))

	for i, ch := range dests {
		go func(i int, ch <-chan int) {
			defer wg.Done()
			for n := range ch {
				fmt.Printf("get value: %d from ch #%d\n", n, i)
			}
		}(i, ch)
	}
	wg.Wait()
}
