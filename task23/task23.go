package task23

import (
	"fmt"
)

func deleteFromSlice[T any](s *[]T, index int) {
	*s = append((*s)[:index], (*s)[index+1:]...)
}

func Run() {
	fmt.Println("--- Task 23 ---")

	a := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(a)

	deleteFromSlice(&a, 3)

	fmt.Println(a)

}
