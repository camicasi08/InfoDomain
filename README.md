# InfoDomain

Este es un proyecto para consultar la información del protocolo SSL de un dominio específico.

## Tecnologías
El proyecto está construido con las siguientes tecnologías.En el backend se utilizó: Golang, y CockroachDB
Para el Frontend la herramienta principal fue  Vuejs
###
### Instalación y puesta en marcha

Instalar dependencias Go
```sh
$ go get -u github.com/lib/pq
$ go get -u github.com/valyala/fasthttp
$ go get -u github.com/buaazp/fasthttprouter
$ go get -u github.com/PuerkitoBio/goquery
$ go get -ugithub.com/likexian/whois-go
```
Ejecución del servidor http en GO. El servidor escucha sobre el puerto 3000
```sh
$ go run main.go
```
Dependencias Vuejs

```sh
$ cd front
$ npm install -d
$ npm run serve
```
### Consumo del API
##### Endpoints

##### /info/:domain
Retorna la información del dominio consutlado en la siguiente estructura JSON:
```json
{
    “servers”: [
        {
            “address”: “server1”,
            “ssl_grade”: “B”,
            “country”: “US”,
            “owner”: “Amazon.com, Inc.”
        },
        {
            “address”: “server2”,
            “ssl_grade”: “A+”,
            “country”: “US”
            “owner”: “Amazon.com, Inc.”
        },
        {
            “address”: “server3”,
            “ssl_grade”: “A”,
            “country”: “US”
            “owner”: “Amazon.com, Inc.”
        }
    ],
    “servers_changed”: true,
    “ssl_grade”: “B”,
    “previous_ssl_grade”: “A+”,
    “logo”: “ https://server.com/icon.png ”,
    “title”: “Title of the page”,
    “is_down”: false
}
```
##### /recents
Retorna la información de los dominios consultados recientemente.

