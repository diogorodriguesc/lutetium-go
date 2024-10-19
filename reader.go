package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func (sitemap *Sitemap) Location() string { return sitemap.xml.location }

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

func main() {
	sitemap := &Sitemap{
		xml: Xml{
			location: "sitemaps.xml",
		},
	}

	_, err := fmt.Fprintf(os.Stdout, "Location %s\n", sitemap.Location())
	if err != nil {
		return
	}

	UrlSet := sitemap.Read()

	fmt.Fprintf(os.Stdout, "Sitemap Xmlns %v\n", UrlSet.Xmlns)
}
