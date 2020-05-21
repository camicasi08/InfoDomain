

package main

import (
	//"./services"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"./controllers"
	
	"log"
)


func main() {
	
	router := fasthttprouter.New()

	router.GET("/info/:domain", controllers.GetInfoDomain)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
	//services.GetDomain("rappi.com")
	//services.Whois("45.5.164.11")
	//fmt.Println(string(result))
}
