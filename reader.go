package lutetiumgo

import (
	"encoding/xml"
	"fmt"
	"os"
)

func (sitemap *Sitemap) Location() string { return sitemap.Xml.Location }

func (sitemap *Sitemap) Read() *UrlSet {
	file, err := os.Open(sitemap.Location())
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	var UrlSet UrlSet

	err = decoder.Decode(&UrlSet)
	if err != nil {
		fmt.Printf("Error decoding XML: %v\n", err)
		return nil
	}

	return &UrlSet
}
