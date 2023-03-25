package task16

import (
	"fmt"
)

func quicksort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	p := s[0]
	less := make([]int, 0, len(s))
	greater := make([]int, 0, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] <= p {
			less = append(less, s[i])
		}
		if s[i] > p {
			greater = append(greater, s[i])
		}
	}
	less = quicksort(less)
	greater = quicksort(greater)
	res := append(less, p)
	res = append(res, greater...)
	return res
}

func Run() {
	fmt.Println("--- Task 16 ---")

	nums := []int{8, 54, 3, 0, -1, -245, 43, 1}

	nums = quicksort(nums)

	fmt.Println(nums)
}
