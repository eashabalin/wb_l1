package task11

import (
	"fmt"
	"wb_l1/set"
)

func Run() {
	fmt.Println("--- Task 11 ---")

	set1 := set.NewSet([]int{1, 2, 3, 3, 4, 4, 5, 5, 5})
	set2 := set.NewSet([]int{2, 2, 2, 3, 5, 6, 6})

	res := set1.Intersection(set2)

	res.Print()
}
