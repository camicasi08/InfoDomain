package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"reflect"
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
