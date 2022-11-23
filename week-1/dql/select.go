package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID      int    `sql:"id"`
	Name    string `sql:"name"`
	Age     int    `sql:"age"`
	Address string `sql:"address"`
	Salary  int    `sql:"salary"`
}

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

	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		panic(err)
	}

	var listEmployee []Employee

	// melakukan looping untuk menampung data dari rows ke struct Employee
	for rows.Next() {
		var employee Employee

		err = rows.Scan(&employee.ID, &employee.Name, &employee.Age, &employee.Address, &employee.Salary)
		if err != nil {
			panic(err)
		}

		listEmployee = append(listEmployee, employee)
	}

	fmt.Println("Berhasil melakukan query tabel employee")
	fmt.Println(listEmployee)

}
