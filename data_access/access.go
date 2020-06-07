package data_access

import (
	"database/sql"
	"fmt"
	"log"

	"../models"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/Info_Domain?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	return db
}

func GetDomains(db *sql.DB) []models.Domain {
	query := `SELECT id, d.name,
	d.ssl_grade,
	d.previous_ssl_grade,
	d.servers_changed,
	d.is_down, logo, title FROM domain d ORDER BY updated DESC;`
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
		domain.Servers = getServersByDomain(db, domain.Id)
		domains = append(domains, domain)
		fmt.Printf("%d %s %s %s %t %t %s %s\n", id, name, sslGrade, previousSslGrade, serversChanged, isDown, logo, title)
	}
	//fmt.Println(domains)
	return domains
}

func AddDomain(db *sql.DB, domain models.Domain) {

	insertedID := 0
	query := fmt.Sprintf(`INSERT INTO public."domain" (id, "name", ssl_grade, previous_ssl_grade, is_down, servers_changed, title, logo, created, updated)
	VALUES(unique_rowid(), '%s', '%s', '%s', %t, %t, '%s', '%s', now(), now()) RETURNING id;`, domain.Name, domain.Ssl_grade, domain.Previous_ssl_grade, domain.Is_down, domain.Servers_changed, domain.Title, domain.Logo)
	err := db.QueryRow(query).Scan(&insertedID)
	fmt.Printf("ID: %d", insertedID)
	if err == nil {
		fmt.Println("Creado")
		for _, elem := range domain.Servers {
			fmt.Println(elem)
			AddServer(db, elem, insertedID)
		}
	}
	/*
		fmt.Print(query)
		if _, err := db.Exec(
			query); err != nil {
			log.Fatal(err)
		} else {

		} */
}

func getServersByDomain(db *sql.DB, idDomain int) []models.Server {
	var servers []models.Server
	query := fmt.Sprintf(`SELECT id, address, ssl_grade, owner, country FROM server WHERE id_domain = %d`, idDomain)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var server models.Server
		var id int
		var sslGrade, address, owner, country string
		if err := rows.Scan(&id, &address, &sslGrade, &owner, &country); err != nil {
			log.Fatal(err)
		}

		//server.Id = id
		server.Address = address
		server.Ssl_grade = sslGrade
		server.Owner = owner
		server.Country = country

		servers = append(servers, server)
	}
	fmt.Println(servers)
	return servers

}

func AddServer(db *sql.DB, server models.Server, idDomain int) {
	query := fmt.Sprintf(`INSERT INTO public."server"
		(id, id_domain, address, owner, country, ssl_grade, created, updated)
		VALUES(unique_rowid(), %d, '%s', '%s', '%s', '%s', now(), now()); `,
		idDomain, server.Address, server.Owner, server.Country, server.Ssl_grade)

	//fmt.Print(query)
	if _, err := db.Exec(
		query); err != nil {
		log.Fatal(err)
	}
}

func FindDomainByName(db *sql.DB, nameDomain string) models.Domain {
	query := fmt.Sprintf(`SELECT id, d.name,
	d.ssl_grade,
	d.previous_ssl_grade,
	d.servers_changed,
	d.is_down, logo, title FROM domain d WHERE name = '%s';`, nameDomain)
	var id int
	var name string
	var sslGrade string
	var previousSslGrade string
	var serversChanged bool
	var isDown bool
	var logo string
	var title string
	//var servers []models.Server

	var domain models.Domain

	row := db.QueryRow(query)
	switch err := row.Scan(&id, &name, &sslGrade, &previousSslGrade, &serversChanged, &isDown, &logo, &title); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		domain.Id = id
		domain.Name = name
		domain.Ssl_grade = sslGrade
		domain.Previous_ssl_grade = previousSslGrade
		domain.Servers_changed = serversChanged
		domain.Is_down = isDown
		domain.Logo = logo
		domain.Title = title
		domain.Servers = getServersByDomain(db, domain.Id)
	default:
		panic(err)

	}

	return domain

}

func UpdateDomain(db *sql.DB, domain models.Domain, idDomain int) {
	sqlStatement := fmt.Sprintf(`
		UPDATE domain
		SET ssl_grade = '%s',previous_ssl_grade = '%s', name = '%s', logo = '%s', title = '%s', is_down = %t, servers_changed = %t, updated = now()
		WHERE id = %d;`, domain.Ssl_grade, domain.Previous_ssl_grade, domain.Name, domain.Logo, domain.Title, domain.Is_down, domain.Servers_changed, idDomain)
	res, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

func DeleteServers(db *sql.DB, idDomain int) {
	sqlStatement := fmt.Sprintf(`DELETE FROM server	WHERE id_domain = %d;`, idDomain)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

/* func main() {
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
} */
