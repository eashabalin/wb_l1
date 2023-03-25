package task7

import (
	"fmt"
	"sync"
)

func writeDataToSyncMap(m *sync.Map) {
	m.Store("one", 1)
	fmt.Println("Written to map: m[\"one\"] = 1")
}

func readDataFromSyncMap(m *sync.Map) {
	val, ok := m.Load("one")
	if ok {
		fmt.Println("Read from map: m[\"one\"] = ", val)
	} else {
		fmt.Println("No value")
	}
}

func Run2() {
	var m sync.Map

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		writeDataToSyncMap(&m)
		wg.Done()
	}()

	go func() {
		readDataFromSyncMap(&m)
		wg.Done()
	}()

	wg.Wait()

}
