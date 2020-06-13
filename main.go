package main

import (
	//"./services"
	"./controllers"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"log"
)

var (
	corsAllowHeaders     = "authorization"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	}
}
func main() {

	router := fasthttprouter.New()

	router.GET("/info/:domain", controllers.GetInfoDomain)
	router.GET("/recents", controllers.GetRecentDomains)
	log.Fatal(fasthttp.ListenAndServe(":3000", CORS(router.Handler)))

}
