package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"../models"
	"github.com/likexian/whois-go"
)

var grades = map[string]int{
	"Current": 11,
	"A+":      10,
	"A":       9,
	"A-":      8,
	"B":       7,
	"C":       6,
	"D":       5,
	"E":       4,
	"F":       3,
	"M":       2,
	"T":       1,
}

//GetDomain : Funciòn para obtener información desde el API SSL
func GetDomain(domain string) (models.Domain, error) {
	var respuesta models.Domain
	var error1 error
	response1, err1 := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)
	if err1 != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err1)
		error1 = errors.New("The HTTP request failed with errors")
		respuesta.Is_down = true
	} else {
		fmt.Println(response1.Body)
		response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)

		respuesta.Ssl_grade = "Current"
		respuesta.Name = domain
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			error1 = errors.New("The HTTP request failed with errors")
			respuesta.Is_down = true
		} else {

			defer response.Body.Close()
			data, _ := ioutil.ReadAll(response.Body)
			jsonResult := string(data)
			var result map[string]interface{}
			json.Unmarshal([]byte(jsonResult), &result)
			endpoints, ok := result["endpoints"].([]interface{})
			if !ok {
				fmt.Println("Error obteniendo endpoints")
			}
			if result["status"] == "ERROR" || result["status"] == "DNS" {
				respuesta.Is_down = true
			}

			var servers []models.Server

			//Access to endpoints

			if endpoints != nil {
				for _, m := range endpoints {
					endpoint := m.(map[string]interface{})
					var server models.Server
					grade := endpoint["grade"]
					ipAddress := endpoint["ipAddress"]
					host := endpoint["serverName"]
					if grade != nil {
						server.Ssl_grade = grade.(string)
						if grades[server.Ssl_grade] < grades[respuesta.Ssl_grade] {
							respuesta.Ssl_grade = server.Ssl_grade
						}
					}

					if ipAddress != nil {
						server.Address = ipAddress.(string)
						organization, country := Whois(server.Address)
						server.Owner = organization
						server.Country = country

					} else if host != nil {
						server.Address = host.(string)
					}

					servers = append(servers, server)

				}
				//fmt.Println("INFO SERVERS")
				//fmt.Println(servers)
			}
			respuesta.Title, respuesta.Logo = Scraper(domain)
			respuesta.Servers = servers

		}
		if respuesta.Ssl_grade == "Current" {
			respuesta.Ssl_grade = ""
		}
		//fmt.Println("RESPONSE")
		//fmt.Println(respuesta)
	}

	if error1 != nil {
		return respuesta, error1
	} else {
		return respuesta, nil
	}

}

//Whois : Implementación del método Whois para obtener información de los servidores
func Whois(ip string) (string, string) {
	whoisRaw, err := whois.Whois(ip)
	organization := ""
	country := ""
	if err == nil {
		//fmt.Println(whois_raw)
		organization = getOrganization(whoisRaw)
		country = getCountry(whoisRaw)
		//fmt.Println(organization)
		//fmt.Println(country)
	}
	return organization, country
}

func getOrganization(whoisRaw string) string {

	r := regexp.MustCompile("owner.+|orgname.+|org-name.+|organization.+")
	organization := ""
	if r.MatchString(strings.ToLower(whoisRaw)) {
		organization = r.FindString(strings.ToLower(whoisRaw))
		array := strings.Split(organization, ":")
		organization = strings.Trim(array[1], " ")
		return organization

	}
	return organization

}

func getCountry(whoisRaw string) string {

	r := regexp.MustCompile("country.+")
	country := ""
	if r.MatchString(strings.ToLower(whoisRaw)) {
		country = r.FindString(strings.ToLower(whoisRaw))
		array := strings.Split(country, ":")
		country = strings.Trim(array[1], " ")
		if len(country) > 2 {
			//fmt.Println("country 1")
			sub := []rune(country)
			newcountry := string(sub[0:3])
			newcountry = strings.Trim(newcountry, " ")
			//fmt.Println(newcountry)
			country = newcountry
			//country = strings.Trim(array[1], " ")

		}
		return country

	}
	return country

}
