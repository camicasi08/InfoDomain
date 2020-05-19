package main

import (
	"./services"
)

func main() {
	
	//services.GetDomain("www.facebook.com")
	services.Whois("1.2.1.1")
	//fmt.Println(string(result))
}
