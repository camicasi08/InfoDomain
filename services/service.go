package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getServer(domain string) {
	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func main() {
	getServer("www.miputo.com")
}
