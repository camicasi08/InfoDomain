package services

import (
	"fmt"
	//"log"
	"errors"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Scraper(domain string) (string, string, error) {

	response, err := http.Get("http://www." + domain)
	var error1 error
	var logos []string
	if err != nil {
		fmt.Print(err)
		error1 = errors.New("The HTTP request failed with errors")

	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err == nil && error1 == nil {
		fmt.Print(err)
		error1 = errors.New("The HTTP request failed with errors")
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

		return title, logo, nil

	} else {
		error1 = errors.New("Scraping failed")
		return "", "", error1
	}

}
