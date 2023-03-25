package task20

import (
	"fmt"
	"strings"
)

func Run() {
	fmt.Println("--- Task 20 ---")

	str := "snow dog sun"

	words := strings.Split(str, " ")

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	fmt.Println(words)
}
