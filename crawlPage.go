package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s':%v\n", rawCurrentURL, err)
		return
	}

	//skip other websites
	if currentURL.Hostname() != cfg.baseURL.Host {
		return
	}

	normURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("couldn't normalize current URL: %v\n", err)
		return
	}

	isFirst := cfg.addPageVsited(normURL)
	if !isFirst {
		return
	}

	fmt.Printf("Crawling page - %s\n", normURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("coudln't get html from '%v': %v\n", normURL, err)
		return
	}

	urls, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		fmt.Printf("Couldn't parse HTML for '%v': %v\n", normURL, err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		cfg.crawlPage(url)
	}
}
