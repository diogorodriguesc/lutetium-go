package main

import (
	"fmt"
	"os"

	lutetiumgo "github.com/diogorodriguesc/lutetium-go"
)

func main() {
	sitemap := &lutetiumgo.Sitemap{
		Xml: lutetiumgo.Xml{
			Location: "../../sitemaps.xml",
		},
	}

	_, err := fmt.Fprintf(os.Stdout, "Location %s\n", sitemap.Location())
	if err != nil {
		return
	}

	UrlSet := sitemap.Read()

	fmt.Fprintf(os.Stdout, "Sitemap Xmlns %v\n", UrlSet.Xmlns)
}
