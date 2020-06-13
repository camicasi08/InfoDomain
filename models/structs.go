package models

//Server struct
import "time"

type Server struct {
	Address   string `json:"address"`
	Ssl_grade string `json:"ssl_grade"`
	Country   string `json:"country"`
	Owner     string `json:"owner"`
	Id        int    `json:"-"`
}

//Response struct
type Domain struct {
	Servers            []Server  `json:"servers"`
	Servers_changed    bool      `json:"servers_changed"`
	Ssl_grade          string    `json:"ssl_grade"`
	Previous_ssl_grade string    `json:"previous_ssl_grade"`
	Logo               string    `json:"logo"`
	Title              string    `json:"title"`
	Is_down            bool      `json:"is_down"`
	Id                 int       `json:"-"`
	Name               string    `json:"name"`
	Updated            time.Time `json:"-"`
}
