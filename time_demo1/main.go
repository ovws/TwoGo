package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(time.Now())
	// year := time.Now().Year()
	// month := time.Now().Month()
	// day := time.Now().Day()
	// hour := time.Now().Hour()
	// minute := time.Now().Minute()
	// second := time.Now().Second()
	// fmt.Printf("现在是：%d年%d月%d日%d时%d分%d秒", year, month, day, hour, minute, second)
	// timeStamp1 := time.Now().Unix()
	// timeStamp2 := time.Now().UnixNano()
	// fmt.Println()
	// fmt.Println(timeStamp1)
	// fmt.Println(timeStamp2)

	// // 通过时间戳获取时间
	// timeObj := time.Unix(timeStamp1, 0)
	// fmt.Println(timeObj)

	// n := 3
	// time.Sleep(time.Duration(n) * time.Second)

	now := time.Now()
	fmt.Println(now)
	t2 := now.Add(time.Hour * 1)
	fmt.Println(t2)
	// for tmp := range time.Tick(time.Millisecond * 3) {
	// 	fmt.Println(tmp)
	// }
	// ret1 := now.Format("2006-01-02 03:04:05 PM")
	// fmt.Println(ret1)

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// ret2 := now.In(loc).Format("2006-01-02 03:04:05 PM")
	// fmt.Println(ret2)
	timeStr := "2021--01--01 00:00:00"
	timeObj, err := time.Parse("2006--01--02 15:04:05", timeStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	timeObj2, err := time.ParseInLocation("2006--01--02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj2)

}
