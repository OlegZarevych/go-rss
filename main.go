package main

import (
	"fmt"
	reader "go-rss/internal/reader"
)

func main() {
	rss := reader.ReadRss()
	fmt.Println(rss)
}
