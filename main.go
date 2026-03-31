package main

import (
	"fmt"
	reader "go-rss/internal/reader"
	xml "go-rss/internal/xml"
)

// main is the entry point of the application. It fetches the RSS feed, prints it as a
// raw string, and then parses it into an XML structure before printing the result.
func main() {
	rssAsString := reader.ReadRss()
	fmt.Println(rssAsString)
	fmt.Println("---------")
	fmt.Println("Add XML struct support")
	rssAsBytes := reader.ReadRssBytes()
	rssAsXml := xml.NewXml(rssAsBytes)
	fmt.Println(rssAsXml)
}
