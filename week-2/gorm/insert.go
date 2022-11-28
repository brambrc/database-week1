package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	City      string
	gorm.Model
}

func main() {
	dns := "host=localhost user=postgres password=Kmzway87a@ dbname=learningOMR port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var employees = []Employee{
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "mail1@mail.com",
			City:      "New York",
		},
		{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "mail2@mail.com",
			City:      "New York",
		},
		{
			FirstName: "John",
			LastName:  "Smith",
			Email:     "mail3@mail.com",
			City:      "Jakarta",
		},
	}
	result := db.Create(&employees)

	if result.Error != nil {
		panic(err)
	}
	fmt.Println("Insert Data Successfull")

}
