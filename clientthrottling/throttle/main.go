package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Effector[R, W any] func(context.Context, *R) (*W, error)

type token struct{}

func Throttle[R, W any](ctx context.Context, effector Effector[R, W], max uint, refill uint, d time.Duration) Effector[R, W] {
	tokens := make(chan token, max)
	ticker := time.NewTicker(d)

	go func() {
		for {
			select {
			case <-ticker.C:
				for i := uint(0); i < refill; i++ {
					tokens <- token{}
				}
			case <-ctx.Done():
				fmt.Println("token refilling done")
				return
			}
		}
	}()

	return func(ctx context.Context, r *R) (*W, error) {
		select {
		case <-tokens:
			return effector(ctx, r)
		case <-ctx.Done():
			fmt.Println("effetor stop")
			return nil, ctx.Err()
		}
	}
}

func myServiceStub(ctx context.Context, s *string) (*string, error) {
	t := strings.ToUpper(*s)
	return &t, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	tStub := Throttle(ctx, myServiceStub, 100, 10, time.Duration(1000000000))
	req := "abcd"
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done in main")
			return
		default:
			time.Sleep(time.Duration(rand.Intn(3)*100) * time.Microsecond)
			res, err := tStub(ctx, &req)
			if err != nil {
				log.Printf("big : %v\n", err)
			} else {
				log.Printf("res is : %v\n", *res)
			}
		}
	}
}
