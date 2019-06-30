package sitemapparser

import (
	"testing"
)

func TestSitemapIndexVersion(t *testing.T) {
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

	t.Error("fail")
}

func TestSitemapVersion(t *testing.T) {
	t.Error("fail")
}

func TestGzippedVersion(t *testing.T) {
	t.Error("fail")
}

func TestNotGzippedVersion(t *testing.T) {
	t.Error("fail")
}
