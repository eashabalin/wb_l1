package task15

import (
	"fmt"
)

func createHugeString() string {
	s := "string"
	for i := 0; i < 10; i++ {
		s += s
	}
	return s
}

// Убираем глобальную переменную, убираем лишную переменную в функции someFunc() для хранения огромной строки

func someFunc() string {
	return createHugeString()[:100]
}

func Run() {
	fmt.Println("--- Task 15 ---")

	fmt.Println(someFunc())

}
