package lutetiumgo

import "encoding/xml"

type Xml struct {
	Location string
}

type Sitemap struct {
	Xml Xml
}

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}
