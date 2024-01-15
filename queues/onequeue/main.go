package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"

	"github.com/tbian7/cloud/queues/onequeue/impl"
)

func randYield() {
	if rand.Intn(2) == 0 {
		runtime.Gosched()
	}
}

func randCancel() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	if rand.Intn(4) == 0 {
		cancel()
	} else {
		_ = cancel
	}
	return ctx
}

func main() {
	q := impl.NewQueue()

	var wg sync.WaitGroup
	for n := 20; n > 0; n-- {
		wg.Add(1)
		go func(n int) {
			randYield()
			if item, err := q.Get(randCancel()); err != nil {
				fmt.Printf("tianbian %v\n", err)
			} else {
				fmt.Printf("%2d: %2d\n", n, item)
			}
			wg.Done()
		}(n)
	}

	for i := 0; i < 100; i++ {
		q.Put(i)
		randYield()
	}

	wg.Wait()
}
