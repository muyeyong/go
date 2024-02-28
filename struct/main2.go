package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 4}
	p := &v
	p.X = 2
	fmt.Println(v)
}