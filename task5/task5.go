package task5

import (
	"fmt"
	"time"
)

func Publish(ch chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
	close(ch)
}

func Read(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func Run() {
	fmt.Println("--- Task 5 ---")

	const timeout = 5

	// Создаём таймер, по истечению которого программа завершается
	timer := time.NewTimer(time.Second * timeout)

	ch := make(chan int)

	// Публикуем данные в канал
	go Publish(ch)

	// Читаем данные из канала
	go Read(ch)

	// Ожидаем получения сигнала от таймера. После программа завершается.
	<-timer.C

}
