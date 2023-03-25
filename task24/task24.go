package task24

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p1 *Point) distance(p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func Run() {
	fmt.Println("--- Task 24 ---")

	p1 := NewPoint(12.2, 14.5)
	p2 := NewPoint(11.3, 5.4)

	fmt.Println(p1.distance(p2))

}
