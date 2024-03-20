package main

import "fmt"

func do(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func main() {
	c := make(chan int)
	arr := []int{1, 4, -1, 4, 6, 0}
	go do(arr[:len(arr)/2], c)
	go do(arr[len(arr)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}