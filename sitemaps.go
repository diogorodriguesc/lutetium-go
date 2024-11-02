package lutetiumgo

import (
	"encoding/xml"
	"fmt"
	"os"
)

func (sitemap *Sitemap) Read() *UrlSet {
	file, err := os.Open(sitemap.Path)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	decoder := xml.NewDecoder(file)

	var UrlSet UrlSet

	err = decoder.Decode(&UrlSet)
	if err != nil {
		fmt.Printf("Error decoding XML: %v\n", err)
		return nil
	}

	return &UrlSet
}
