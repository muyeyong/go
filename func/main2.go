package main

import (
	"fmt"
	"math"
)

type V struct {
	X, Y float64
}

func (v *V) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *V) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := V{3, 4}
	fmt.Println(v.Abs())
	v.Scale(5)
	fmt.Println(v.Abs())
}