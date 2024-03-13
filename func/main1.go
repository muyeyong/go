package main

import (
	"fmt"
	"math"
)

type V struct {
	X, Y float64
}

func (v V) Abs() float64 {
	return math.Abs(v.X * v.X + v.Y * v.Y)
}

func FuncAbs(v V) float64 {
	return math.Abs(v.X * v.X + v.Y * v.Y)
} 

func main () {
	v := V{3, 4}
	fmt.Println(v.Abs())
	m := &v
	fmt.Println(m)
	fmt.Println(FuncAbs(*m))
}