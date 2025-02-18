package main

import "fmt"

func main() {
	var x interface{}
	x = 100

	if x.(int) == 100 {
		fmt.Println("x is int")
	} else {
		fmt.Println("x is not int")
	}
}
