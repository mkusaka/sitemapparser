package sitemapparser

import (
	"bytes"
	"compress/gzip"
	"errors"
	"log"
	"net/http"

	"github.com/beevik/etree"
)

// @return sitemapped url which all
func Scheduler(url string) ([]string, error) {
	// download and parse xml
	sitemapXML, err := Downloader(url)

	if err != nil {
		log.Fatal(err)
	}

	parsedURLs, isSitemap, err := Parser(sitemapXML)

	if err != nil {
		log.Fatal(err)
	}

	parsedSiteURLs := []string{}
	if isSitemap {
		for _, indexURL := range parsedURLs {
			parsedURLs2, err := Scheduler(indexURL)
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

func Downloader(url string) (string, error) {
	log.Printf("start download: %s", url)

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

	out := bytes.Buffer{}
	out.ReadFrom(reader)
	s := string(out.Bytes())

	return s, nil
}

// @str downloaded sitemap
// @return
func Parser(sitemapXML string) ([]string, bool, error) {
	isSitemapIndex := false
	xmlStr := sitemapXML
	doc := etree.NewDocument()
	if err := doc.ReadFromString(xmlStr); err != nil {
		log.Fatal(err)
		return nil, isSitemapIndex, err
	}
	sitemapSet := doc.SelectElement("sitemapindex")
	urlSet := doc.SelectElement("urlset")
	siteUrls := []string{}
	sitemaps := []*etree.Element{}
	if urlSet != nil {
		sitemaps = urlSet.SelectElements("url")
	} else if sitemapSet != nil {
		isSitemapIndex = true
		sitemaps = sitemapSet.SelectElements("sitemap")
	} else {
		return []string{}, false, errors.New("something wrong string: " + sitemapXML)
	}
	for _, sitemap := range sitemaps {
		loc := sitemap.SelectElement("loc")
		siteUrls = append(siteUrls, loc.Text())
	}

	return siteUrls, isSitemapIndex, nil
}
