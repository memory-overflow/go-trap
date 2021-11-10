package main

import "fmt"
import "time"
import "context"

func main() {
	tick1 := time.NewTicker(1*time.Second)
	tick2 := time.NewTicker(2*time.Second)
	tick3 := time.NewTicker(4*time.Second)
	ctx := context.Background()
	fmt.Println("Start:\t",
	time.Now().Format("2006-01-02 15:04:05"))
	LOOP: for {
		select {
		case <-ctx.Done():
			return
		case <-tick3.C:
			fmt.Printf("[%s] do %d 秒定时任务\n",
				time.Now().Format("2006-01-02 15:04:05"), 4)
			break LOOP
		case <-tick2.C:
			fmt.Printf("[%s] do %d 秒定时任务\n",
				time.Now().Format("2006-01-02 15:04:05"), 2)
		case <-tick1.C:
			fmt.Printf("[%s] do %d 秒定时任务\n",
				time.Now().Format("2006-01-02 15:04:05"), 1)
		}
		fmt.Println("")
	}
}
