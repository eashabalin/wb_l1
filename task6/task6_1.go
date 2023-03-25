package task6

import (
	"fmt"
	"time"
)

func Worker1(done <-chan bool, stopped chan<- bool) {
	for i := 1; true; i++ {
		select {
		case <-done:
			time.Sleep(time.Second)
			fmt.Println("Worker", 1, "stopped...")
			stopped <- true
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Worker", 1, "working:", i, "...")
		}
	}

}

func Run1() {
	fmt.Println("--- Task 6.1 ---")

	done := make(chan bool)
	stopped := make(chan bool)

	go Worker1(done, stopped)

	time.Sleep(time.Second * 3)

	done <- true
	<-stopped

}
