package lutetiumgo

import (
	"testing"
)

func TestSitemap_Creation(t *testing.T) {
	sitemap := &Sitemap{
		Path: "https://location/sitemap.xml",
	}

	if sitemap.Path != "https://location/sitemap.xml" {
		t.Errorf("Expected is %s, got %s", "https://location/sitemap.xml", sitemap.Path)
	}
}

func TestSitemap_Read(t *testing.T) {
	sitemap := &Sitemap{
		Path: "sitemaps.xml",
	}

	UrlSet := sitemap.Read()

	if UrlSet.Xmlns != "http://www.sitemaps.org/schemas/sitemap/0.9" {
		t.Errorf("Expected is %s, got %s", "http://www.sitemaps.org/schemas/sitemap/0.9", UrlSet.Xmlns)
	}

	if len(UrlSet.Urls) != 3 {
		t.Errorf("Expected UrlSet.Urls count is 3, got %d", len(UrlSet.Urls))
	}

	var ExpectedUrls = [3]string{"https://www.example.com/a", "https://www.example.com/b", "https://www.example.com/c"}
	var ExpectedLastMods = [3]string{"2024-10-19", "2024-10-18", "2024-10-17"}

	i := 0
	for _, Urls := range UrlSet.Urls {
		if Urls.Loc != ExpectedUrls[i] {
			t.Errorf("Expected is %s, got %s", ExpectedUrls[i], Urls.Loc)
		}
		if Urls.LastMod != ExpectedLastMods[i] {
			t.Errorf("Expected is %s, got %s", ExpectedLastMods[i], Urls.LastMod)
		}
		i++
	}
}
