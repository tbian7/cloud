package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn((1000))) * time.Millisecond)
		}
	}()
	return c
}
