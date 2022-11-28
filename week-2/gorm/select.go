package main

import (
	"encoding/json"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID        int    `sql:"id" gorm:"primaryKey"`
	FirstName string `sql:"first_name"`
	LastName  string `sql:"last_name"`
	Email     string `sql:"email"`
	City      string `sql:"city"`
	gorm.Model
}

func main() {
	dns := "host=localhost user=postgres password=Kmzway87a@ dbname=learningOMR port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	employees := []Employee{}

	rows := db.Find(&employees)

	if err != nil {
		panic(err)
	}
	if rows != nil {
		employeesJSON, err := json.Marshal(employees)
		if err != nil {
			panic(err)
		}
		println(string(employeesJSON))
	}

}
