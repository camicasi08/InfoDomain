package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"reflect"
	"github.com/likexian/whois-go"
	//"github.com/likexian/whois-parser-go"
	"regexp"
	"strings"

)



func GetDomain(domain string) {
	//response1, err1:= http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)
	//fmt.Println(response1, err1)
	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		jsonResult := string(data)
		//fmt.Println(jsonResult)
		var result map[string]interface{}
		json.Unmarshal([]byte(jsonResult), &result)
		endpoints, ok := result["endpoints"].([]interface{})
		if !ok {
			fmt.Println("Error obteniendo endpoints")
		}

		fmt.Println(reflect.TypeOf(result))
		fmt.Println(reflect.TypeOf(endpoints))
		fmt.Println(endpoints)
		

		//Access to endpoints 
		
		if endpoints != nil {
			for _, m := range endpoints {
				endpoint := m.(map[string]interface{})
				grade  := endpoint["grade"]
				ipAddress := endpoint["ipAddress"]
				host := endpoint["serverName"]
				fmt.Println(grade, ipAddress, host) 
			}
		}
	
	}
}

func Whois(ip string){
	whois_raw, err := whois.Whois(ip)
	if err == nil {
		fmt.Println(whois_raw)
		name, extension := searchDomain(whois_raw)
		fmt.Println(name)
		fmt.Println(extension)
		/* result, err2 := whoisparser.ParseContact(whois_raw)
		if err2 == nil{
			fmt.Println(result.Domain)
		}else
		{
			fmt.Println(err2)
		} */
	}
}


func searchDomain(text string) (string, string) {
	r := regexp.MustCompile(`(?i)\[?domain(\s*\_?name)?\]?[\s\.]*\:?\s*([a-z0-9\-\.]+)\.([a-z]{2,})`)
	m := r.FindStringSubmatch(text)
	if len(m) > 0 {
		return strings.ToLower(strings.TrimSpace(m[2])), strings.ToLower(strings.TrimSpace(m[3]))
	}

	r = regexp.MustCompile(`(?i)\[?domain(\s*\_?name)?\]?\s*\:?\s*([a-z]{2,})\n`)
	m = r.FindStringSubmatch(text)
	if len(m) > 0 {
		return strings.ToLower(strings.TrimSpace(m[2])), ""
	}

	return "", ""
}