package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"
)

func test(ctx context.Context) error {
	time.Sleep(20 * time.Millisecond)
	return errors.New("")
}

func main() {
	times := 1000
	for i := 0; i < times; i++ {
		ctx, _ := context.WithTimeout(
			context.Background(), 10*time.Millisecond)
		done := make(chan error, 1)
		go func() {
			done <- test(ctx)
		}()
		select {
		case <-done:
			break
		case <-ctx.Done():
			break
		}
	}
	time.Sleep(1 * time.Second)
	fmt.Println("runtime number", runtime.NumGoroutine())
}
