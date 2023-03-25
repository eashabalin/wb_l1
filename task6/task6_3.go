package task6

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	stop chan bool
}

func NewWorker() *Worker {
	return &Worker{stop: make(chan bool)}
}

func (w *Worker) Run() {
	for i := 1; true; i++ {
		select {
		default:
			time.Sleep(time.Second)
			fmt.Println("Worker", 3, "working:", i, "...")
		case <-w.stop:
			return
		}
	}
}

func (w *Worker) Stop() {
	w.stop <- true
	fmt.Println("Worker 3 stopped")
}

func Run3() {
	fmt.Println("--- Task 6.3 ---")

	var wg sync.WaitGroup
	w := NewWorker()

	wg.Add(1)
	go func() {
		w.Run()
		wg.Done()
	}()

	time.Sleep(time.Second * 3)

	w.Stop()

	wg.Wait()
}
