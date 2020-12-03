package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 544
	user     = "postgres"
	password = "secret"
	dbname   = "go_web_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected")
	defer db.Close()
}
