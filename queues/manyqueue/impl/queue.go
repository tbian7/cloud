package impl

type waiter[T any] struct {
	n int
	r chan []T
}

type state[T any] struct {
	is []T
	ws []waiter[T]
}

type queue[T any] struct {
	s chan state[T]
}

func NewQueue[T any]() *queue[T] {
	s := make(chan state[T], 1)
	s <- state[T]{}
	return &queue[T]{s: s}
}

func (q *queue[T]) Put(item T) {
	s := <-q.s
	s.is = append(s.is, item)
	for len(s.ws) > 0 {
		w := s.ws[0]
		if len(s.is) < w.n {
			break
		}
		w.r <- s.is[:w.n:w.n]
		close(w.r)
		s.is = s.is[w.n:]
		s.ws = s.ws[1:]

	}
	q.s <- s

}

func (q *queue[T]) GetMany(n int) []T {
	s := <-q.s
	if len(s.ws) == 0 && len(s.is) >= n {
		items := s.is[:n:n]
		s.is = s.is[n:]
		q.s <- s
		return items
	}

	r := make(chan []T)
	w := waiter[T]{
		n: n,
		r: r,
	}
	s.ws = append(s.ws, w)
	q.s <- s
	return <-r
}
