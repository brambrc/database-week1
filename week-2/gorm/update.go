package main

import (
	"fmt"

	_ "github.com/lib/pq"
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

	result := db.Model(&Employee{}).Where("id = ?", 2).Update("city", "Surabaya")

	if result.Error != nil {
		panic(err)
	}

	fmt.Println("Update Data Successfull")

}
