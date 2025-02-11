package main

import (
	"fmt"

	"github.com/wikiq/TwoGo/infra"
)

func getRetriever() retriever {
	return infra.Retriever{}
}

type retriever interface {
	Get(url string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://ifconfig.me"))
}
