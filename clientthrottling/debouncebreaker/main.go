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

func Breaker[R, W any](circuit Circuit[R, W], failureThreadhold uint64) Circuit[R, W] {
	lastStateSuccessful := true
	var consecutiveFailures uint64
	lastAttempt := time.Now()

	return func(ctx context.Context, r *R) (*W, error) {
		if consecutiveFailures >= failureThreadhold {
			backoffLevel := consecutiveFailures - failureThreadhold
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << time.Duration(backoffLevel))

			if !time.Now().After(shouldRetryAt) {
				return nil, errors.New("circuit open -- service unreacheable")
			}
		}

		lastAttempt = time.Now()
		w, err := circuit(ctx, r)
		if err != nil {
			if !lastStateSuccessful {
				consecutiveFailures++
			}
			lastStateSuccessful = false
			return w, err
		}
		lastStateSuccessful = true
		consecutiveFailures = 0
		return w, err
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
	bStub := DebounceFirst(Breaker(myServiceStub, 2), 10*time.Millisecond)
	req := "abcd"

	for {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		res, err := bStub(ctx, &req)
		if err != nil {
			if err == ErrDebounced {
				log.Printf("it is %v.\n", err)
				continue
			}
			log.Printf("fail to call my service: %v\n", err)
			continue
		}
		log.Println(*res)
	}
}
