package task2

import (
	"fmt"
	"sync"
)

func PrintSquared(num int) {
	fmt.Println(num * num)
}

func Run() {
	fmt.Println("--- Task 2 ---")

	nums := [...]int{2, 4, 6, 8, 10}

	// Дождаться выполнения нескольких горутин можно с помощью WaitGroup
	fmt.Println("WaitGroup:")

	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			PrintSquared(num)
			wg.Done()
		}(num)
	}
	wg.Wait()

	// Или же с помощью канала
	fmt.Println("Channel:")

	// Буферизированный канал эффективнее в данном случае, так как он не блокирует выполнение горутин пока не заполнится
	ch := make(chan bool, len(nums))

	for _, num := range nums {
		go func(num int) {
			PrintSquared(num)
			ch <- true
		}(num)
	}

	for range nums {
		<-ch
	}
}
