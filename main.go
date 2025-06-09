package main

import (
	"fmt"
	reader "go-rss/internal/reader"
	xml "go-rss/internal/xml"
)

func main() {
	rssAsString := reader.ReadRss()
	fmt.Println(rssAsString)
	fmt.Println("---------")
	fmt.Println("Add XML struct support")
	rssAsBytes := reader.ReadRssBytes()
	rssAsXml := xml.NewXml(rssAsBytes)
	fmt.Println(rssAsXml)
}
