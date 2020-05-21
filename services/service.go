package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	//"reflect"
	"github.com/likexian/whois-go"
	//"github.com/likexian/whois-parser-go"
	"regexp"
	"strings"
	"../models"

)


var grades = map[string]int{
	"Current":11,
	"A+": 10,
	"A": 9,
	"A-": 8,
	"B": 7,
	"C": 6,
	"D": 5,
	"E": 4,
	"F": 3,
	"M": 2,
	"T": 1,
}



func GetDomain(domain string) models.Response {
	//response1, err1:= http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)
	//fmt.Println(response1, err1)
	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)
	var respuesta models.Response
	respuesta.Ssl_grade = "Current"
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		respuesta.Is_down= true
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

		//fmt.Println(reflect.TypeOf(result))
		//fmt.Println(result)
		//fmt.Println(reflect.TypeOf(endpoints))
		//fmt.Println(endpoints)
		//fmt.Println(result["status"])
		if result["status"] == "ERROR" || result["status"] == "DNS"{
			respuesta.Is_down = true
		}
	
		var servers [] models.Server


		//Access to endpoints 
		
		if endpoints != nil {
			for _, m := range endpoints {
				endpoint := m.(map[string]interface{})
				var server models.Server
				grade  := endpoint["grade"]
				ipAddress := endpoint["ipAddress"]
				host := endpoint["serverName"]
				if grade != nil{
					server.Ssl_grade = grade.(string)
					if grades[server.Ssl_grade] < grades[respuesta.Ssl_grade]{
						respuesta.Ssl_grade = server.Ssl_grade
					}
				}

				if ipAddress != nil {
					server.Address = ipAddress.(string)
					organization, country := Whois(server.Address)
					server.Owner = organization
					server.Country = country

				}else if host != nil {
					server.Address = host.(string)
				}
			
				servers = append(servers, server)

			}
			//fmt.Println("INFO SERVERS")
			//fmt.Println(servers)
		}
		respuesta.Logo, respuesta.Title = Scraper(domain)
		respuesta.Servers = servers
	
	
	}
	if respuesta.Ssl_grade == "Current"{
		respuesta.Ssl_grade = ""
	}
	//fmt.Println("RESPONSE")
	fmt.Println(respuesta)
	return respuesta
}

func Whois(ip string)(string, string) {
	whois_raw, err := whois.Whois(ip)
	organization:= ""
	country := ""
	if err == nil {
		//fmt.Println(whois_raw)
		organization = getOrganization(whois_raw)
		country = getCountry(whois_raw)
		//fmt.Println(organization)
		//fmt.Println(country)
	}
	return organization, country
}

func  getOrganization(whois_raw string) string{
	
	r := regexp.MustCompile("owner.+|orgname.+|org-name.+|organization.+")
	organization := ""
	if r.MatchString(strings.ToLower(whois_raw)) {
		organization = r.FindString(strings.ToLower(whois_raw))
		array := strings.Split(organization, ":")
		organization =strings.Trim(array[1], " ")
		return organization
		
	}
	return organization

}

func  getCountry(whois_raw string) string{
	
	r := regexp.MustCompile("country.+")
	country := ""
	if r.MatchString(strings.ToLower(whois_raw)) {
		country = r.FindString(strings.ToLower(whois_raw))
		array := strings.Split(country, ":")
		country =strings.Trim(array[1], " ")
		return country
		
	}
	return country

}
