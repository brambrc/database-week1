package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	dns := "host=localhost user=postgres password=Kmzway87a@ dbname=company2 port=5432 sslmode=disable"

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

	_, err = db.Exec(`CREATE TABLE departments (
		id INT PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	)`)

	if err != nil {
		panic(err)
	}

	// create table employee
	_, err = db.Exec(`CREATE TABLE employees (
		id INT PRIMARY KEY,
		id_department INT NOT NULL REFERENCES departments(id), 
		name VARCHAR(255) NOT NULL,
		age INT NOT NULL,
		address VARCHAR(255),
		salary INT
	  )`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO departments
      VALUES 	(1, 'IT'),
				(2, 'Marketing'),
				(3, 'Finance')`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO employees
      VALUES 	(1, 1,'Rizki', 25, 'Jl. Kebon Jeruk', 2000000),
				(2, 1,'Andi', 27, 'Jl. Kebon Sirih', 3000000),
				(3, 2,'Budi', 30, 'Jl. Kebon Melati', 4000000),
				(4, 2,'Caca', 32, 'Jl. Kebon Anggrek', 5000000),
				(5, 3,'Deni', 35, 'Jl. Kebon Mawar', 6000000)`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Table employees & department created")
}
