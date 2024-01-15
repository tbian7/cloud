package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"runtime"
	"sync"
)

var (
	addr     net.Addr
	addrOnce sync.Once
)

func dial() (net.Conn, error) {
	addrOnce.Do(func() {
		ln, err := net.Listen("tcp", ":0")
		if err != nil {
			panic(err)
		}
		addr = ln.Addr()
		go func() {
			for {
				in, err := ln.Accept()
				if err != nil {
					return
				}
				go io.Copy(os.Stdout, in)
			}
		}()
	})
	return net.Dial(addr.Network(), addr.String())
}

type Pool struct {
	sem  chan token
	idle chan net.Conn
}
type token struct{}

func NewPool(limit int) *Pool {
	return &Pool{
		sem:  make(chan token, limit),
		idle: make(chan net.Conn, limit),
	}
}

func (p *Pool) Acquire(
	ctx context.Context) (
	net.Conn, error) {

	select {
	case conn := <-p.idle:
		return conn, nil
	case p.sem <- token{}:
		conn, err := dial()
		if err != nil {
			<-p.sem
		}
		return conn, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (p *Pool) Release(c net.Conn) {
	p.idle <- c
}

func (p *Pool) Hijack(c net.Conn) {
	<-p.sem
}

// hacky.
func (p *Pool) Close() {
	close(p.idle)
	for i := range p.idle {
		i.Close()
	}
}

func randYield() {
	if rand.Intn(4) == 0 {
		runtime.Gosched()
	}
}

func main() {
	p := NewPool(3)

	var wg sync.WaitGroup
	for n := 10; n > 0; n-- {
		wg.Add(1)
		go func(n int) {
			randYield()
			defer wg.Done()

			conn, err := p.Acquire(context.Background())
			if err != nil {
				panic(err)
			}
			defer p.Release(conn)
			fmt.Println(n)
			fmt.Fprintf(conn, "Hello from goroutine %d on connection %p!\n", n, conn)
			fmt.Println("abc", n)
			runtime.Gosched()
		}(n)
	}

	for n := 4; n > 0; n-- {
		wg.Add(1)
		go func() {
			randYield()
			defer wg.Done()

			conn, err := p.Acquire(context.Background())
			if err != nil {
				panic(err)
			}
			defer p.Hijack(conn)

			fmt.Fprintf(conn, "Goodbye from hijacked connection %p!\n", conn)

			runtime.Gosched()
		}()
	}

	wg.Wait()
	// A hacky way to fix the bug: program exit before connection is closed (buffer flushed)
	// p.Close()
}
