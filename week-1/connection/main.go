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
	password = "Kmzway87a@"
	dbname   = "testDatabase"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	return db, nil
}

func main() {
	_, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
