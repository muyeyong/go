package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type){
	case int:
		fmt.Printf("%v is int\n", v)
	case string:
		fmt.Printf("%q is string\n", v)
	default:
		fmt.Printf("%v is unknown\n", v)
	}
}

func main() {
	do(1)
	do("hello")
	do(false)
}