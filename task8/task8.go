package task8

import "fmt"

func setBit(num int64, pos uint) int64 {
	num = num | (1 << pos)
	return num
}

func clearBit(num int64, pos uint) int64 {
	var mask int64 = ^(1 << pos)
	num = num & mask
	return num
}

func Run() {
	fmt.Println("--- Task 8 ---")

	var num int64 = 123524

	fmt.Printf("%b\n", num)

	num = setBit(num, 0)

	num = clearBit(num, 2)
	
	fmt.Printf("%b\n", num)

}
