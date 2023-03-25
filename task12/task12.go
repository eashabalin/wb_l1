package task12

import (
	"fmt"
	"wb_l1/set"
)

func Run() {
	fmt.Println("--- Task 12 ---")

	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	s := set.NewSet(strings)

	s.Print()
}
