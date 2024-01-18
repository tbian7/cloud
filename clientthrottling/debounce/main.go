package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Circuit[R, W any] func(context.Context, *R) (*W, error)

var ErrDebounced = errors.New("debounced")

func DebounceFirst[R, W any](circuit Circuit[R, W], d time.Duration) Circuit[R, W] {
	var threshold time.Time = time.Now()

	return func(ctx context.Context, r *R) (*W, error) {
		var w *W
		err := ErrDebounced
		if threshold.Before(time.Now()) {
			w, err = circuit(ctx, r)
		}
		threshold = time.Now().Add(d)
		return w, err
	}
}

func myServiceStub(ctx context.Context, s *string) (*string, error) {

	t := strings.ToUpper(*s)
	return &t, nil
}

func main() {
	ctx := context.Background()
	dStub := DebounceFirst(myServiceStub, 10*time.Millisecond)
	req := "abcd"
	time.Sleep(time.Second)
	for {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		res, err := dStub(ctx, &req)
		if err != nil {
			if err == ErrDebounced {
				log.Printf("it is %v.\n", err)
				continue
			}
			log.Fatalf("big : %v\n", err)
		}
		log.Printf("res is : %v\n", *res)
	}
}
