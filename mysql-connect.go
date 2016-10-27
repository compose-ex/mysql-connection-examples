package main

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

func main() {

	db, err := sql.Open("mysql", "admin:mypass@tcp(aws-us-east-1-portal.23.dblayer.com:15942)/compose?tls=skip-verify")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		log.Fatal(err)
	}

	var dbNames string

	for rows.Next() {
		err := rows.Scan(&dbNames)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dbNames)
	}
}
