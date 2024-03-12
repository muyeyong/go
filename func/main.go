package main

import (
	"fmt"
	"math"
)

type VT struct {
	X, Y float64
}

func (v VT) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func Abs2(v VT) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := VT{3, 4}
	fmt.Println(v.Abs(), Abs2(v))
}