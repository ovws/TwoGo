// Package main 实现了一个简单的网页内容下载器，支持HTTP/HTTPS协议
package main

import (
	"fmt"
	"log"

	"github.com/wikiq/TwoGo/infra"
)

// URLRetriever 定义了一个获取网页内容的接口
type URLRetriever interface {
	// Get 获取指定URL的内容
	// url: 要获取内容的URL地址
	// 返回获取到的网页内容
	Get(url string) string
}

// newRetriever 创建并返回一个实现了URLRetriever接口的实例
func newRetriever() URLRetriever {
	return infra.Retriever{}
}

func main() {
	retriever := newRetriever()
	url := "https://ifconfig.me"

	content := retriever.Get(url)
	if content == "" {
		log.Printf("Failed to retrieve content from %s", url)
		return
	}

	fmt.Println(content)
}
