package main

import "fmt"

func main() {
	var x interface{}
	x = 100
	getMath, k := x.(int)
	if k {
		fmt.Println(getMath)
	}
	// if x.(int) == 100, k {
	// 	fmt.Println("x is int")
	// } else {
	// 	fmt.Println("x is not int")
	// }
}
