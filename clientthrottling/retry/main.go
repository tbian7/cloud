package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Effector[R, W any] func(context.Context, *R) (*W, error)

func Retry[R, W any](effector Effector[R, W], retries int, delay time.Duration) Effector[R, W] {
	return func(ctx context.Context, r *R) (*W, error) {
		for i := 0; ; i++ {
			w, err := effector(ctx, r)
			if err == nil || i >= retries {
				return w, err
			}
			log.Printf("Attemp %d failed; retrying in %v\n", i+1, delay)

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
	}
}

func myServiceStub(ctx context.Context, s *string) (*string, error) {
	if rand.Intn(10) < 5 {
		return nil, errors.New("service unavaliable")
	}

	t := strings.ToUpper(*s)
	return &t, nil
}

func main() {
	ctx := context.Background()
	rStub := Retry(myServiceStub, 2, time.Second)
	req := "abcd"
	res, err := rStub(ctx, &req)
	if err != nil {
		log.Printf("got err in %v\n", err)
	} else {
		log.Printf("res is : %v\n", *res)
	}
}
