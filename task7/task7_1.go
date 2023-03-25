package task7

import (
	"fmt"
	"sync"
)

func writeData(m map[string]int, mutex *sync.RWMutex) {
	mutex.Lock()
	defer mutex.Unlock()
	m["one"] = 1
	fmt.Println("Written to map: m[\"one\"] = 1")
}

func readData(m map[string]int, mutex *sync.RWMutex) {
	mutex.RLock()
	defer mutex.RUnlock()
	fmt.Println("Read from map: m[\"one\"] = ", m["one"])
}

func Run1() {
	fmt.Println("--- Task 7 ---")

	m := make(map[string]int)
	var mutex sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		writeData(m, &mutex)
		wg.Done()
	}()

	go func() {
		readData(m, &mutex)
		wg.Done()
	}()

	wg.Wait()
}
