package main

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

func (p *Point) CalcTo(to *Point) float64 {
	// Немножко математики - такое меня устраивает :)
	return math.Sqrt(math.Pow(math.Abs(p.x-to.x), 2) + math.Pow(math.Abs(p.y-to.y), 2))
}

func main() {
	a := NewPoint(2, 3)
	b := NewPoint(0, 5)

	fmt.Println(a.CalcTo(b))
}
