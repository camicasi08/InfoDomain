package services


import (
    //"fmt"
    "log"
    "net/http"
	"github.com/PuerkitoBio/goquery"
	"strings"
)


func Scraper(domain string)(string, string){

	response, err := http.Get("http://www."+domain)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}

	var logos []string
	doc.Find("head").Each(func(i int, s *goquery.Selection) {
		s.Find("link").Each(func(j int, l *goquery.Selection){
			//fmt.Println(l)
			
			band, ok := l.Attr("href")
			if ok {
				if strings.Contains(band, "jpg") || strings.Contains(band, "png") || strings.Contains(band, "jpge"){
					logos = append(logos,band)
				}
				
			}
			
		})
		
	})

	title := doc.Find("title").Text()
	var logo string
	if len(logos) > 0{
		logo = logos[0]
	}
	return title, logo
	
}