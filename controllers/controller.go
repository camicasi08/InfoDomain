package controllers

import (
	"encoding/json"
	"fmt"

	"strings"

	"../data_access"
	"../services"
	"github.com/valyala/fasthttp"
)

//GetInfoDomain : Se obtiene la información de dominio y sus servidores
func GetInfoDomain(ctx *fasthttp.RequestCtx) {
	domain := ctx.UserValue("domain").(string)
	domain = strings.ToLower(domain)
	result := services.GetDomain(ctx.UserValue("domain").(string))
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		ctx.Write(b)
		conn := data_access.Connect()

		resultDB := data_access.FindDomainByName(conn, domain)
		fmt.Println(resultDB)

		if resultDB.Id != 0 {
			fmt.Println("SI HAY RESULTADOS")
			data_access.UpdateDomain(conn, result, resultDB.Id)
			data_access.DeleteServers(conn, resultDB.Id)
			for _, elem := range result.Servers {
				data_access.AddServer(conn, elem, resultDB.Id)
			}

		} else {
			fmt.Println("NO HAY RESULTADOS")
			data_access.AddDomain(conn, result)
		}

		/* 	if !reflect.DeepEqual(resultDB.Servers, result.Servers) {
			fmt.Println("DIFF")
		} */

		//fmt.Fprint(ctx, string(b))
	}

}

//GetRecentDomains : Se obtienen los dominios consultados previamente
func GetRecentDomains(ctx *fasthttp.RequestCtx) {
	conn := data_access.Connect()
	domains := data_access.GetDomains(conn)
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	b, err := json.Marshal(domains)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		ctx.Write(b)
		//fmt.Fprint(ctx, string(b))
	}

}
