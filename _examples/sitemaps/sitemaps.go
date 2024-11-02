package main

import (
	"fmt"
	"os"

	lutetiumgo "github.com/diogorodriguesc/lutetium-go"
)

func main() {
	sitemap := &lutetiumgo.Sitemap{
		Path: "../../sitemaps.xml",
	}

	fmt.Fprintf(os.Stdout, "Path %s\n", sitemap.Path)

	UrlSet := sitemap.Read()

	fmt.Fprintf(os.Stdout, "Sitemap Xmlns %v\n", UrlSet.Xmlns)
}
