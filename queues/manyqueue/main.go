package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"

	"github.com/tbian7/cloud/queues/manyqueue/impl"
)

func main() {
	q := impl.NewQueue[int]()

	var wg sync.WaitGroup
	for n := 10; n > 0; n-- {
		wg.Add(1)
		go func(n int) {
			if rand.Intn(3) == 0 {
				runtime.Gosched()
			}
			items := q.GetMany(n)
			fmt.Printf("%2d: %2d\n", n, items)
			wg.Done()
		}(n)
	}

	for i := 0; i < 100; i++ {
		q.Put(i)
		runtime.Gosched()
	}

	wg.Wait()

}
