package sitemapparser

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/beevik/etree"
)

// @return sitemapped url which all
func scheduler(url string) ([]string, error) {
	// download and parse xml
	sitemapXML, err := downloader(url)

	if err != nil {
		log.Fatal(err)
	}

	parsedURLs, isSitemap, err := parser(sitemapXML)

	if err != nil {
		log.Fatal(err)
	}

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
	fmt.Println(url)

	client := new(http.Client)

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	response, err := client.Do(request)

	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatal(err)
		return "", err
	}
	defer response.Body.Close()

	// unzip
	reader, err := gzip.NewReader(response.Body)
	if err != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatal(err)
		return "", err
	}
	defer reader.Close()
	_, errs := io.Copy(os.Stdout, reader)

	if errs != nil {
		log.SetFlags(log.Lshortfile)
		log.Fatal(err)
		return "", err
	}

	out := bytes.Buffer{}
	out.ReadFrom(reader)
	s := string(out.Bytes())

	return s, nil
}

// @str downloaded sitemap
// @return
func parser(sitemapXML string) ([]string, bool, error) {
	isSitemapIndex := false
	xmlStr := sitemapXML
	doc := etree.NewDocument()
	if err := doc.ReadFromString(xmlStr); err != nil {
		log.Fatal(err)
		return nil, isSitemapIndex, err
	}
	set := doc.SelectElement("sitemapindex")
	// if index is nil(not sitemap index) then parse
	siteUrls := []string{}
	sitemaps := []*etree.Element{}
	if set == nil {
		set = doc.SelectElement("urlset")
		sitemaps = set.SelectElements("url")
	} else {
		isSitemapIndex = true
		sitemaps = set.SelectElements("sitemap")
	}
	for _, sitemap := range sitemaps {
		loc := sitemap.SelectElement("loc")
		siteUrls = append(siteUrls, loc.Text())
	}

	return siteUrls, isSitemapIndex, nil
}
