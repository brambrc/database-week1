package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	dns := "host=localhost user=postgres password=Kmzway87a@ dbname=testDatabase port=5432 sslmode=disable"

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	// create table employee
	_, err = db.Exec(`CREATE TABLE employee (
		id INT,
		name VARCHAR(255) NOT NULL,
		age INT NOT NULL,
		address VARCHAR(255),
		salary INT
	  )`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Table employee created")
}
