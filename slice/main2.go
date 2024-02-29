package main

import "fmt"

func main() {
	v := []string {"1", "8", "9"}
	fmt.Println((v))
	v1 := []struct{
		x int
		y string
	}{
		{1, "2"},
	}
	fmt.Println(v1)
}