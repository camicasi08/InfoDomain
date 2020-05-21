package controllers

import (
	"../services"
	"fmt"
	"github.com/valyala/fasthttp"
	"encoding/json"
)


func GetInfoDomain(ctx *fasthttp.RequestCtx){
	result := services.GetDomain(ctx.UserValue("domain").(string))
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	b, err := json.Marshal(result)
    if err != nil {
        fmt.Println(err)
        return
    }else{
		ctx.Write(b)
		//fmt.Fprint(ctx, string(b))
	}
   
	
}