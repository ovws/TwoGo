package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(time.Now())
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	fmt.Printf("现在是：%d年%d月%d日%d时%d分%d秒", year, month, day, hour, minute, second)
	timeStamp1 := time.Now().Unix()
	timeStamp2 := time.Now().UnixNano()
	fmt.Println()
	fmt.Println(timeStamp1)
	fmt.Println(timeStamp2)

	// 通过时间戳获取时间
	timeObj := time.Unix(timeStamp1, 0)
	fmt.Println(timeObj)

	n := 3
	time.Sleep(time.Duration(n) * time.Second)
}
