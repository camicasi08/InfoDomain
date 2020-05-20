package main

import (
	"./services"
)

func main() {
	
	services.GetDomain("www.unal.edu.co")
	//services.Whois("45.5.164.11")
	//fmt.Println(string(result))
}
