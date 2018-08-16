package main

import (
	"fmt"
	//"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

type Entry struct {
	gorm.Model
	Day  string
	Data string
}

func setupDB() *gorm.DB {
	db, err := gorm.Open(
		"postgres",
		"host=localhost port=5432 user=six dbname=book password=dorjeChupa",
	)
	if err != nil {
		panic("Could not connect to the database.")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Entry{})

	user := &User{Name: "admin", Password: "ballet"}
	if db.NewRecord(user) {
		db.Create(user)
	}

	entries := []Entry{}
	db.Find(&entries)
	difference := 3 - len(entries)
	for i := 0; i < difference; i++ {
		entry := &Entry{
			Day:  fmt.Sprintf("2018.01.0%d", (i + 1)),
			Data: "{\"vow\": \"Protect life\", \"content\": \"some explanation...\"}",
		}
		db.Create(entry)
	}

	return db
}
