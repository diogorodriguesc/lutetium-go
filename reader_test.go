package lutetiumgo

import "testing"

func TestSitemap_Location(t *testing.T) {
	sitemap := &Sitemap{
		Xml: Xml{
			Location: "https://location/sitemap.xml",
		},
	}

	if sitemap.Location() != "https://location/sitemap.xml" {
		t.Errorf("Expected is %s, got %s", "https://location/sitemap.xml", sitemap.Location())
	}
}

func TestSitemap_Read(t *testing.T) {
	sitemap := &Sitemap{
		Xml: Xml{
			Location: "sitemaps.xml",
		},
	}

	UrlSet := sitemap.Read()

	if UrlSet.Xmlns != "http://www.sitemaps.org/schemas/sitemap/0.9" {
		t.Errorf("Expected is %s, got %s", "http://www.sitemaps.org/schemas/sitemap/0.9", UrlSet.Xmlns)
	}

	// test is incomplete
}
