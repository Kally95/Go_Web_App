package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "go_web_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var id int
	var name, email string
	// The first version w/out the ID
	rows, err := db.Query(`
	 SELECT id, name, email
	 FROM users`)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			panic(err)
		}
		fmt.Println("id", id, "name", name, "email", email)
	}
}
