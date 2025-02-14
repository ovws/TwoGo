package main

import (
	"fmt"

	"github.com/wikiq/TwoGo/infra/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://ifconfig.me")

}

func main() {
	var r Retriever
	r = mock.Retriever{"Hello World"}
	fmt.Println(download(r))
	// fmt.Println(download(mock.Retriever{"Hello World"}))
}
