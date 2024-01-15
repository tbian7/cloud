package main

import "fmt"

type token struct{}
type task struct {
	id int
}

const limit = 10

func main() {
	var tasks []task
	for i := 0; i < 1000; i++ {
		tasks = append(tasks, task{id: i})
	}

	sem := make(chan token, limit)

	for _, t := range tasks {
		sem <- token{}
		go func(t task) {
			fmt.Println(t)
			<-sem
		}(t)
	}

	// barrier
	for n := limit; n > 0; n-- {
		sem <- token{}
	}
}
