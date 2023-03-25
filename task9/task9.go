package task9

import (
	"fmt"
	"sync"
)

func Run() {
	fmt.Println("--- Task 9 ---")

	nums := [...]int{1, 3, 5, 7, 9, 11, 13, 15}

	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		for _, n := range nums {
			ch1 <- n
		}
		close(ch1)
		wg.Done()
	}()

	go func() {
		for n := range ch1 {
			ch2 <- n * 2
		}
		close(ch2)
		wg.Done()
	}()

	go func() {
		for n := range ch2 {
			fmt.Println(n)
		}
		wg.Done()
	}()

	wg.Wait()

}
