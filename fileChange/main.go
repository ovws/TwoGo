package main

import (
	"fmt"
	"os"
)

func readFromFile() {
	fileObj, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer fileObj.Close()
	var tmp = make([]byte, 128)
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	fmt.Println(n)
	fmt.Println(string(tmp[:n]))
}

func readall() {
	fileObj, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer fileObj.Close()
	var tmp = make([]byte, 128)
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	fmt.Println(n)
	fmt.Println(string(tmp[:n]))
}

func main() {

}
