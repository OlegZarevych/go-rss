package main

import (
	"fmt"
	db "go-rss/internal/db"
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

	conn, err := db.ConnectDb()
	if err != nil {
		panic(err)
	}
	fmt.Println(conn)
	db.CreateTable(conn)

	for _, item := range rssAsXml.Channel.Items {
		fmt.Println(item)
		id, err := db.CreateRecord(conn, item)
		if err != nil {
			panic(err)
		}
		fmt.Println("Record ID: $1", id)
	}

}
