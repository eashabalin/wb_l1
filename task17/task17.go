package task17

import "fmt"

func binarySearch(s []int, n int) (int, bool) {
	left := 0
	right := len(s) - 1
	for left <= right {
		mid := (left + right) / 2
		guess := s[mid]
		if guess == n {
			return mid, true
		} else if guess > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0, false
}

func Run() {
	fmt.Println("--- Task 17 ---")

	nums := []int{1, 2, 4, 6, 10, 12, 14, 16, 17, 18, 24, 55, 78, 79, 89, 90, 91, 99, 100, 52423, 43435464}

	index, found := binarySearch(nums, 43435464)
	fmt.Println(index, found)

	index, found = binarySearch(nums, 0)
	fmt.Println(index, found)

	index, found = binarySearch(nums, 18)
	fmt.Println(index, found)
}
