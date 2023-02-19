package main

import (
	"fmt"

	colly "github.com/gocolly/colly/v2"
)

var DOMAINS_FOR_SCAPING = []string{
	"https://www.caranddriver.com/",
	"https://www.autonews.com/",
	"https://www.topgear.com/",
}

var URL_FOR_SCRAPING = []string{
	"https://www.caranddriver.com/news/",
	"https://www.autonews.com/news",
	"https://www.topgear.com/car-news",
}

var USER_AGENTS = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
}

func main() {
	newsCollector := colly.NewCollector(
		colly.UserAgent(USER_AGENTS[0]),
	)

	pageCollector := newsCollector.Clone()

	isFirst := true
	newsCollector.OnHTML("a[data-vars-cta=\"4 Across Block 0\"]", func(e *colly.HTMLElement) {
		if isFirst {
			fmt.Println(e.Request.AbsoluteURL(e.Attr("href")))
			err := pageCollector.Visit(e.Request.AbsoluteURL(e.Attr("href")))
			isFirst = false
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	})

	pageCollector.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println("Main text:", e.Text)
	})

	pageCollector.OnHTML("p", func(e *colly.HTMLElement) {
		if isFirst {
			fmt.Println("Additional text:", e.Text)
			isFirst = false
		}
	})

	newsCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := newsCollector.Visit(URL_FOR_SCRAPING[0])
	if err != nil {
		fmt.Println("Can't visit the page")
		return
	}
}
