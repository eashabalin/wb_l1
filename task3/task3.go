package task3

import "fmt"

// Будем считать квадраты в одельных горутинах. По готовности результаты будут прибавляться к значению в главной горутине.

func Squared(num int) int {
	return num * num
}

func Run() {
	fmt.Println("--- Task 3 ---")

	nums := [...]int{2, 4, 6, 8, 10}

	ch := make(chan int, len(nums))

	for _, num := range nums {
		go func(num int) {
			ch <- Squared(num)
		}(num)
	}

	res := 0

	for range nums {
		res += <-ch
	}

	fmt.Printf("Сумма квадратов %v равна %d\n", nums, res)

}
