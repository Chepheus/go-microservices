package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/Chepheus/go-microservices/web-scraper-service/site_scraper"
	colly "github.com/gocolly/colly/v2"
)

var USER_AGENTS = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
}

func main() {
	urlsForScrap := getUrlsForScrap()
	for _, scrapVars := range urlsForScrap {
		randomNumber := getRandInRange(0, len(USER_AGENTS))

		newsCollector := colly.NewCollector(
			colly.UserAgent(USER_AGENTS[randomNumber]),
		)

		scrapResponse := site_scraper.ScrapData(newsCollector, scrapVars)
		str, _ := json.MarshalIndent(scrapResponse, "", "\t")
		fmt.Println(string(str))
	}
}

func getUrlsForScrap() []site_scraper.ScrapingVars {
	var urlsForScraping []site_scraper.ScrapingVars
	urlsForScraping = append(urlsForScraping, site_scraper.NewScrapingVars(
		"https://www.caranddriver.com/news/",
		"a[data-vars-cta=\"4 Across Block 0\"]",
		"h1",
		"p",
	))

	urlsForScraping = append(urlsForScraping, site_scraper.NewScrapingVars(
		"https://www.autonews.com/news",
		".top-stories-image.crain-gallery-node-link-wrapper>a",
		"h1",
		"h2.article-sub-title",
	))

	return urlsForScraping
}

func getRandInRange(min, max int) int {
	return rand.Intn(max-min) + min
}
