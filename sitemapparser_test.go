package sitemapparser

import (
	"testing"
)

func TestSitemapIndexParser(t *testing.T) {
	sitemapIndexXML := `
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/siteindex.xsd" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <sitemap>
    <loc>http://test.me/sitemaps/sitemap1.xml.gz</loc>
    <lastmod>2019-04-30T05:50:07+09:00</lastmod>
  </sitemap>
  <sitemap>
    <loc>http://test.me/sitemaps/sitemap2.xml.gz</loc>
    <lastmod>2019-04-29T01:42:19+09:00</lastmod>
  </sitemap>
  <sitemap>
    <loc>http://test.me/sitemaps/sitemap19.xml.gz</loc>
    <lastmod>2019-04-14T15::42:18+09:00</lastmod>
  </sitemap>
</sitemapindex>`

	urls, isSitemapIndex, err := Parser(sitemapIndexXML)

	if isSitemapIndex == false {
		t.Error("invalid parse: this string is sitemap index.")
	}

	if len(urls) == 0 {
		t.Error("invalid parse: url must be.")
	}

	if err != nil {
		t.Error("invalid parse: error should not to be.")
	}
}

func TestSitemapParser(t *testing.T) {
	sitemapXML := `
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.domain.com /</loc>
    <lastmod>2017-01-01</lastmod>
    <changefreq>weekly</changefreq>
    <priority>0.8</priority>
  </url>
  <url>
    <loc>http://www.domain.com/catalog?item=vacation_hawaii</loc>
    <changefreq>weekly</changefreq>
  </url>
  <url>
    <loc>http://www.domain.com/catalog?item=vacation_new_zealand</loc>
    <lastmod>2008-12-23</lastmod>
    <changefreq>weekly</changefreq>
  </url>
  <url>
    <loc>http://www.domain.com/catalog?item=vacation_newfoundland</loc>
    <lastmod>2008-12-23T18:00:15+00:00</lastmod>
    <priority>0.3</priority>
  </url>
  <url>
    <loc>http://www.domain.com/catalog?item=vacation_usa</loc>
    <lastmod>2008-11-23</lastmod>
  </url>
</urlset>`

	urls, isSitemapIndex, err := Parser(sitemapXML)

	if isSitemapIndex == true {
		t.Error("invalid parse: this string is not sitemap index.")
	}

	if len(urls) == 0 {
		t.Error("invalid parse: url must be.")
	}

	if err != nil {
		t.Error("invalid parse: error should not to be.")
	}
}

/*
func TestSitemapIndexDonwloader(t *testing.T) {
	sitemapXML, _ := downloader("some good url for sitemap index")

	urls, isSitemapIndex, err := parser(sitemapXML)
	fmt.Println(urls)
	fmt.Println(isSitemapIndex)
	fmt.Println(err)
}

func TestSitemapDonwloader(t *testing.T) {
	sitemapXML, _ := downloader("some good url for sitemap index")

	urls, isSitemapIndex, err := parser(sitemapXML)
	fmt.Println(urls)
	fmt.Println(isSitemapIndex)
	fmt.Println(err)
}

func TestScheduler(t *testing.T) {
	sitemapURLs, err := Scheduler("some good url")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(sitemapURLs)
}
*/
