package services

import (
	"fmt"
	//"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Scraper(domain string) (string, string) {

	response, err := http.Get("http://www." + domain)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}

	var logos []string
	doc.Find("head").Each(func(i int, s *goquery.Selection) {
		s.Find("link").Each(func(j int, l *goquery.Selection) {
			//fmt.Println(l)

			band, ok := l.Attr("href")
			if ok {
				if strings.Contains(band, "jpg") || strings.Contains(band, "png") || strings.Contains(band, "jpge") {
					if strings.Contains(band, "http") {
						logos = append(logos, band)
					}

				}

			}

		})

	})

	title := doc.Find("title").Text()
	var logo string
	if len(logos) > 0 {
		logo = logos[0]
	}
	return title, logo

}
