package task19

import (
	"fmt"
)

func reverse(s string) string {
	runes := []rune(s)
	newStr := ""
	for i := len(runes) - 1; i >= 0; i-- {
		newStr += string(runes[i])
	}
	return newStr
}

func Run() {
	fmt.Println("--- Task 19 ---")

	s := "главрыба"

	fmt.Println(reverse(s))
}
