package task6

import (
	"context"
	"fmt"
	"time"
)

func Worker2(ctx context.Context, stopped chan<- bool) {
	for i := 1; true; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Worker", 2, "stopped...")
			stopped <- true
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Worker", 2, "working:", i, "...")
		}
	}
}

func Run2() {
	fmt.Println("--- Task 6.2 ---")

	stopped := make(chan bool)

	ctx, cancel := context.WithCancel(context.Background())

	go Worker2(ctx, stopped)

	time.Sleep(time.Second * 3)

	cancel()
	<-stopped

}
