package task14

import (
	"fmt"
	"reflect"
)

func printType(value any) {
	t := reflect.TypeOf(value)
	fmt.Println(t)
}

func Run() {
	fmt.Println("--- Task 14 ---")

	printType(234)
	printType("Hi!")
	printType(94.42)
	printType(make(chan bool))

}
