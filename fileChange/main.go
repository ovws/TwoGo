package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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
	for {

		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF {
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Println("read file err:", err)
			return
		}
		// fmt.Println(n)
		fmt.Println(string(tmp[:n]))
	}
}

func readByBufio() {
	fileObj, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		read, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("read end")
			return
		}
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Print(read)
	}
}

func readByIoUtil() {
	ioutil.ReadFile("test.txt")
	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	fmt.Println(string(content))
}

func write() {
	fileObj, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file err:", err)
		return

	}
	defer fileObj.Close()
	str := "田东晓最可爱啦！"
	fileObj.Write([]byte(str))
	fileObj.WriteString("hello world\n")
	fileObj.Close()
}

func writeByBufio() {
	fileObj, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("open file err:", err)
		return

	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	str := "田东晓最可爱啦！"
	writer.WriteString(str)
	writer.Flush()
	fileObj.Close()
}

func writeByIoUtil() {
	str := "田东晓同学最乖了啦！"
	err := ioutil.WriteFile("test.txt", []byte(str), 0644)
	// ioutil.WriteFile("test.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func main() {
	// readFromFile()
	// readall()
	// readByBufio()
	// readByIoUtil()
	// write()
	// writeByBufio()
	writeByIoUtil()
}
