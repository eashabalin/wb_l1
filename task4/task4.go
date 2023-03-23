package task4

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Worker – функция, которая будет запускаться в нескольких горутинах
func Worker(id int, jobs <-chan int) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Println("Worker", id, ": ", j)
	}
}

// GenerateJobs – постоянно публикует данные в канал jobs до тех пор, пока не будет получен сигнал в канале done
func GenerateJobs(jobs chan<- int, done <-chan bool) {
	for i := 1; true; i++ {
		select {
		case <-done:
			close(jobs)
			return
		default:
			fmt.Println(i)
			jobs <- rand.Intn(100)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func Run() {
	fmt.Println("--- Task 4 ---")

	// Количество воркеров
	const numWorkers = 5

	// Канал, в который будут приходить данные
	jobs := make(chan int, 100)

	// WaitGroup нужна для того, чтобы программа не завершалась до тех пор, пока все воркеры завершат работу
	wg := sync.WaitGroup{}

	// Запуск воркеров
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			Worker(id, jobs)
		}(w)
	}

	// Канал для завершения публикации данных
	done := make(chan bool, 1)

	// Запуск публикации данных в канал jobs в отдельной горутине
	go GenerateJobs(jobs, done)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание сигнала завершения программы из ОС (Ctrl+C)
	<-sigs

	// Посылаем сигнал завершения публикации в канал jobs
	done <- true

	// Ждём пока все воркеры завершат работу
	wg.Wait()

}
