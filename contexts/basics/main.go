package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Detach returns a context that keeps all the values of its parent context
// but detaches from the cancellation and error handling.
func Detach(ctx context.Context) context.Context {
	return detachedContext{ctx}
}

type detachedContext struct{ parent context.Context }

func (v detachedContext) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (v detachedContext) Done() <-chan struct{}             { return nil }
func (v detachedContext) Err() error                        { return nil }
func (v detachedContext) Value(key interface{}) interface{} { return v.parent.Value(key) }

func mySleepAndTalk(ctx context.Context, d time.Duration, format string, a ...any) {
	select {
	case <-time.After(d):
		fmt.Printf(format, a...)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}

type key int

const sessionKey key = 63

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, sessionKey, "session_id")
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	if _, ok := ctx.Value(64).(string); !ok {
		fmt.Println("type is important")
	}

	if id, ok := ctx.Value(sessionKey).(string); ok {
		mySleepAndTalk(ctx, 5*time.Second, "hello %s world: %s\n", "non-existing", id)
		mySleepAndTalk(Detach(ctx), 5*time.Second, "hello %s world: %s\n", "detached", id)
	}
}
