package task13

import "fmt"

func Run() {
	fmt.Println("--- Task 13 ---")

	a := 1
	b := 2

	fmt.Printf("a = %d; b = %d\n", a, b)

	a, b = b, a

	fmt.Printf("a = %d; b = %d\n", a, b)
}
