package newscrawler

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// collector is crawler
	c := colly.NewCollector(

		colly.AllowedDomains("www.varzesh.com"),
	)

	c.OnHTML(".gs-c-promo-heading__title", func(e *colly.HTMLElement) {
		// print title of news
		fmt.Println("تیتر خبر:", e.Text)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("خطا:", err)
	})

	c.Visit("https://www.varzesh.com/live_corses")
}
