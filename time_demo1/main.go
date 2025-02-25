package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间并展示各种格式
	now := time.Now()
	fmt.Println("Current time (default format):", now)

	// 展示基本时间运算
	oneHourLater := now.Add(time.Hour * 1)
	fmt.Println("Time after 1 hour:", oneHourLater)

	// 时区处理
	shanghaiLoc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Failed to load location:", err)
		return
	}

	// 时间解析示例
	const (
		inputLayout    = "2006-01-02 15:04:05"
		outputLayout   = "2006年01月02日 03:04:05 PM MST"
		nonStandardStr = "2021-01-01 00:00:00"
	)

	// 解析普通时间字符串
	parsedTime, err := time.Parse(inputLayout, nonStandardStr)
	if err != nil {
		fmt.Println("Time parse error:", err)
		return
	}
	fmt.Println("Parsed time (local):", parsedTime.Format(outputLayout))

	// 带时区的解析
	parsedTimeWithLoc, err := time.ParseInLocation(inputLayout, nonStandardStr, shanghaiLoc)
	if err != nil {
		fmt.Println("Time parse with location error:", err)
		return
	}
	fmt.Println("Parsed time (Shanghai):", parsedTimeWithLoc.In(shanghaiLoc).Format(outputLayout))
}
