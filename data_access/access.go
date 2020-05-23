package main

import (
	"database/sql"
	"fmt"
	"log"

	"../models"
	_ "github.com/lib/pq"
)

func connect() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/Info_Domain?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	return db
}

func getDomains(db *sql.DB) []models.Domain {
	query := `SELECT id, d.name,
	d.ssl_grade,
	d.previous_ssl_grade,
	d.servers_changed,
	d.is_down, logo, title FROM domain d;`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var domains []models.Domain
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var sslGrade string
		var previousSslGrade string
		var serversChanged bool
		var isDown bool
		var logo string
		var title string
		if err := rows.Scan(&id, &name, &sslGrade, &previousSslGrade, &serversChanged, &isDown, &logo, &title); err != nil {
			log.Fatal(err)
		}
		var domain models.Domain
		domain.Id = id
		domain.Name = name
		domain.Ssl_grade = sslGrade
		domain.Previous_ssl_grade = previousSslGrade
		domain.Servers_changed = serversChanged
		domain.Is_down = isDown
		domain.Logo = logo
		domain.Title = title
		domains = append(domains, domain)
		fmt.Printf("%d %s %s %s %t %t %s %s\n", id, name, sslGrade, previousSslGrade, serversChanged, isDown, logo, title)
	}
	fmt.Println(domains)
	return domains
}

func addDomain(db *sql.DB, domain models.Domain) {

	query := fmt.Sprintf(`INSERT INTO public."domain" (id, "name", ssl_grade, previous_ssl_grade, is_down, servers_changed, title, logo, created, updated)
	VALUES(unique_rowid(), '%s', '%s', '%s', %t, %t, '%s', '%s', now(), now());`, domain.Name, domain.Ssl_grade, domain.Previous_ssl_grade, domain.Is_down, domain.Servers_changed, domain.Logo, domain.Title)

	fmt.Print(query)
	if _, err := db.Exec(
		query); err != nil {
		log.Fatal(err)
	}
}
func main() {
	db := connect()
	var domain models.Domain
	domain.Name = "rappi.com"
	domain.Ssl_grade = "A"
	domain.Previous_ssl_grade = "B"
	domain.Logo = ""
	domain.Title = ""
	domain.Is_down = false
	domain.Servers_changed = true

	//addDomain(db, domain)
	getDomains(db)
}
