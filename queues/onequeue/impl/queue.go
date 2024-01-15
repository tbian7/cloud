package impl

import (
	"context"
)

type Item = int

type Queue struct {
	items chan []Item // non-empty slices only
	empty chan bool   // holds true if the queue is empty
}

func NewQueue() *Queue {
	items := make(chan []Item, 1)
	empty := make(chan bool, 1)
	empty <- true
	return &Queue{items, empty}
}

func (q *Queue) Get(ctx context.Context) (Item, error) {
	var items []Item

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case items = <-q.items:
	}

	item := items[0]
	if len(items) == 1 {
		q.empty <- true
	} else {
		q.items <- items[1:]
	}
	return item, nil
}

func (q *Queue) Put(item Item) {
	var items []Item
	select {
	case items = <-q.items:
	case <-q.empty:
	}
	items = append(items, item)
	q.items <- items
}
