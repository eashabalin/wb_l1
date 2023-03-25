package task22

import (
	"fmt"
	"math/big"
)

func Run() {
	fmt.Println("--- Task 22 ---")

	a := big.NewInt(1048576) // 2^20
	b := big.NewInt(3000000)

	c := big.NewInt(0)
	c.Mul(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, c.String())

	d := big.NewInt(0)
	d.Div(a, b)
	fmt.Printf("%v / %v = %v\n", a, b, d.String())

	e := big.NewInt(0)
	e.Add(a, b)
	fmt.Printf("%v + %v = %v\n", b, a, e.String())

	f := big.NewInt(0)
	f.Sub(a, b)
	fmt.Printf("%v - %v = %v\n", a, b, f.String())
}
