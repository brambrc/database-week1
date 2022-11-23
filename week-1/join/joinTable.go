package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type EmployeeData struct {
	ID         int    `sql:"employee_id"`
	Name       string `sql:"employee_name"`
	Department string `sql:"department_name"`
	Age        int    `sql:"employee_age"`
}

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

	// join table `customers` and `departments`
	rows, err := db.Query(`
    SELECT
		employees.id AS employee_id,
        employees.name AS employee_name,
		departments.name as department_name
        employees.age AS employee_age,
    FROM employees JOIN departments
    ON employees.id_department = departments.id
    `)

	if err != nil {
		panic(err)
	}

	var allData []EmployeeData

	// iterate over rows
	for rows.Next() {
		var data EmployeeData

		// scan data from rows to struct
		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.Department,
			&data.Age,
		)

		if err != nil {
			panic(err)
		}

		allData = append(allData, data)
	}

	fmt.Println(allData)

}
