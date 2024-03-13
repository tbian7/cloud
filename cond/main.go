package main

import (
	"fmt"
	"sync"
	"time"
)

type donation struct {
	cond    *sync.Cond
	balance int
}

func foo(goal int, d *donation) {
	d.cond.L.Lock()
	for d.balance < goal {
		d.cond.Wait()
	}
	fmt.Printf("reached my goal: %v", goal)
	d.cond.L.Unlock()
}

func main() {
	d := &donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	go foo(10, d)
	go foo(15, d)

	time.Sleep(100 * time.Millisecond)
	d.cond.L.Lock()
	d.balance = 15
	d.cond.L.Unlock()
	d.cond.Signal()
	// d.cond.Broadcast()

	time.Sleep(100 * time.Second)
}
