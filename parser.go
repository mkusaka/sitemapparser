package sitemapparser

import (
	"log"

	"github.com/beevik/etree"
)

// @return sitemapped url which all
func scheduler(url string) ([]string, error) {
	// // downlaod from url and return downloaded url
	sitemapXML, err := downloader(url)
	// // parse and return url sitemap or sitemapped urls flag

	if err != nil {
		log.Fatal(err)
	}

	parsedURLs, isSitemap, err := parser(sitemapXML)
	// // parse sitemap, returns

	if err != nil {
		log.Fatal(err)
	}

	// // if downloaded sitemap is sitemap index, download again with returnd url
	// // recursion is better answer?
	parsedSiteURLs := []string{}
	if isSitemap {
		for _, indexURL := range parsedURLs {
			parsedURLs2, err := scheduler(indexURL)
			if err != nil {
				log.Fatal(err)
			}
			for _, parsedURL := range parsedURLs2 {
				parsedSiteURLs = append(parsedSiteURLs, parsedURL)
			}
		}
	}
	if len(parsedSiteURLs) > 0 {
		parsedURLs = parsedSiteURLs
	}
	return parsedURLs, err
}

func downloader(url string) (string, error) {
	// awsome downloader

	// downalod part

	// if url end with gz, unzip

	// return unziped string
}

// @str downloaded sitemap
// @return
func parser(sitemapXML string) ([]string, bool, error) {
	isSitemapIndex := false
	xmlStr := sitemapXML
	doc := etree.NewDocument()
	if err := doc.ReadFromString(xmlStr); err != nil {
		panic(err)
	}
	index := doc.SelectElement("sitemapindex")
	// if index is nil(not sitemap index) then parse
	/*
		if index != nil {
			isSitemapIndex = true
		}
	*/
	sitemaps := index.SelectElements("sitemap")
	// write not sitemap index version
	siteUrls := []string{}
	for _, sitemap := range sitemaps {
		loc := sitemap.SelectElement("loc")
		siteUrls = append(siteUrls, loc.Text())
	}

	return sitemaps
}
