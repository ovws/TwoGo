package main

import (
	"fmt"
	"time"
)

func main() {

	timeov := time.Now()
	unTime := timeov.Unix()
	fmt.Println("当前时间戳:", unTime)

	unixTime := unTime // Unix时间戳
	// 将Unix时间戳转换为time.Time对象
	t := time.Unix(int64(unixTime), 0)
	var str = t.Format("2006-01-02 15:04:05")
	fmt.Println(str)
}
