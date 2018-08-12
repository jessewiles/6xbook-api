package main

import (
	//"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name     string
	Password string
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

	user := &User{Name: "admin", Password: "ballet"}
	if db.NewRecord(user) {
		db.Create(user)
	}

	return db
}
