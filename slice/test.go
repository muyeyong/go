package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	v := make([][]uint8, dy) 
	for i := range v {
		k := make([]uint8, dx)
		for j := range k{
			k[j] = uint8((i +j ) / 2)
		}
		v[i] = k
	}
	return v
}

func main() {
	pic.Show(Pic)
}
