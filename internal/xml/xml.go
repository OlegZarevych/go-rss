package internal

import (
	"encoding/xml"
)

type Xml struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Language      string `xml:"language"`
	LastBuildDate string `xml:"lastBuildDate"`
	AtomLink      string `xml:"http://www.w3.org/2005/Atom link"`
	Items         []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
}

func NewXml(bytes []byte) Xml {
	var doc Xml
	xml.Unmarshal(bytes, &doc)
	return doc
}
