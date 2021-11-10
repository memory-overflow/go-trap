package main

import "fmt"
import "time"
import "context"

func test(ctx context.Context) {
	// time.Sleep(5 * time.Second)
	x := 0
	for i := 0; i < 100000000000; i++ {
		select {
			case <- ctx.Done():
				fmt.Println("timeout exit, x = ", x)
				return
			default:
				break
		}
		x += 1
	}
	fmt.Println("Done, x = ", x)
}

func main() {
	ctx, _ := context.WithTimeout(
		context.Background(), time.Second*1)
	done := make(chan error, 0)
	go func() {
		test(ctx)
		done <- nil
	}()
	select {
	case err := <- done:
		fmt.Println(err)
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
  time.Sleep(7 * time.Second)
}