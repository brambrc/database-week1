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

	// rename table employee to employees
	_, err = db.Exec(`ALTER TABLE employee RENAME TO employees`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Table employee renamed to employees")
}
