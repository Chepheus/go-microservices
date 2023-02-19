package site_scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type ScrapingVars struct {
	url              string
	newsLinkSelector string
	mainTextSelector string
	subTextSelector  string
}

func NewScrapingVars(url, newsLinkSelector, mainTextSelector, subTextSelector string) ScrapingVars {
	return ScrapingVars{
		url:              url,
		newsLinkSelector: newsLinkSelector,
		mainTextSelector: mainTextSelector,
		subTextSelector:  subTextSelector,
	}
}

type ScrapResponse struct {
	NewsUrl  string
	MainText string
	SubText  string
}

func ScrapData(newsCollector *colly.Collector, vars ScrapingVars) *ScrapResponse {
	scrapResponse := ScrapResponse{}
	pageCollector := newsCollector.Clone()

	isFirst := true
	newsCollector.OnHTML(vars.newsLinkSelector, func(e *colly.HTMLElement) {
		if isFirst {
			newsUrl := e.Request.AbsoluteURL(e.Attr("href"))
			scrapResponse.NewsUrl = newsUrl
			err := pageCollector.Visit(newsUrl)
			isFirst = false
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	})

	pageCollector.OnHTML(vars.mainTextSelector, func(e *colly.HTMLElement) {
		scrapResponse.MainText = e.Text
	})

	pageCollector.OnHTML(vars.subTextSelector, func(e *colly.HTMLElement) {
		if isFirst {
			scrapResponse.SubText = e.Text
			isFirst = false
		}
	})

	newsCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := newsCollector.Visit(vars.url)
	if err != nil {
		fmt.Println("Can't visit the page")
		return nil
	}

	return &scrapResponse
}
