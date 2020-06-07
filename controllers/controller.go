package controllers

import (
	"encoding/json"
	"fmt"
	"reflect"

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

	conn := data_access.Connect()

	resultDB := data_access.FindDomainByName(conn, domain)
	if resultDB.Id != 0 {

		if !reflect.DeepEqual(resultDB.Servers, result.Servers) {
			result.Servers_changed = true
		}

		result.Previous_ssl_grade = resultDB.Ssl_grade
		b, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			ctx.Write(b)
		}
		fmt.Println("SI HAY RESULTADOS")
		data_access.UpdateDomain(conn, result, resultDB.Id)
		data_access.DeleteServers(conn, resultDB.Id)
		for _, elem := range result.Servers {
			data_access.AddServer(conn, elem, resultDB.Id)
		}

	} else {
		fmt.Println("NO HAY RESULTADOS")
		result.Previous_ssl_grade = result.Ssl_grade
		b, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			ctx.Write(b)
		}
		data_access.AddDomain(conn, result)
	}

	/* b, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		ctx.Write(b)
		conn := data_access.Connect()

		resultDB := data_access.FindDomainByName(conn, domain)

		//ctx.Write(b)
		//fmt.Println(resultDB)

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

		if !reflect.DeepEqual(resultDB.Servers, result.Servers) {
			fmt.Println("DIFF")
		} else {
			fmt.Println("ELSE")
		}


	} */

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
