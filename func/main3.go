package main

import (
	"fmt"
	"math"
)

type AT interface {
	abs() float64
}

type V struct {
	X, Y float64
}

func (v *V) abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

type MyFloat float64 

func (f MyFloat) abs() float64 {
	return float64(f)	
}

func main () {
	v := V{3, 4}
	f := MyFloat(math.Sqrt2)
	var a AT
	a = f
	fmt.Println(a.abs())
	a = &v
	fmt.Println(a.abs())
}