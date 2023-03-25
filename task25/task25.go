package task25

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func Run() {
	fmt.Println("--- Task 25 ---")

	fmt.Println("Sleeping...")

	sleep(time.Second * 3)

	fmt.Println("Awake!")

}
