package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Test struct {
	Z int
}

func main() {
	v := Vertex{3,5}
	fmt.Println(v.X)
	v1 := Test{1}
	fmt.Println(v1.Z)
}