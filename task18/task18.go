package task18

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Run() {
	fmt.Println("--- Task 18 ---")

	var c uint64
	var c1 uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				// Используем метод из atomic, чтобы избежать гонки данных
				atomic.AddUint64(&c, 1)
				c1++
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("c (с atomic):", c)
	fmt.Println("c1 (без atomic):", c1)
}
