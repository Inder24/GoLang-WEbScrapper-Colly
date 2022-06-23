package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func CryptoExample() {
	c := colly.NewCollector()

	c.OnHTML("a.cmc-table__column-name--name", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")
}
