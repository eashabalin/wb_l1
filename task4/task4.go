package task4

import (
	"fmt"
	"math/rand"
	"time"
)

func Worker(id int, jobs <-chan int, results chan<- bool) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Println("Worker", id, ": ", j)
		results <- true
	}
}

func GenerateJobs(jobs chan<- int) {
	for {
		jobs <- rand.Intn(100)
		time.Sleep(100 * time.Millisecond)
	}
}

func Run() {
	fmt.Println("--- Task 4 ---")

	const numJobs = 5
	const numWorkers = 5
	jobs := make(chan int, numJobs)
	results := make(chan bool, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go Worker(w, jobs, results)
	}

	GenerateJobs(jobs)

	for {
		<-results
	}
}
