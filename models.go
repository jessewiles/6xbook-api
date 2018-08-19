package main

import (
	"fmt"
	//"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Day struct {
	ID        string    `json:"id" gorm:"primary_key" sql:"type:uuid;default:uuid_generate_v4()"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Entry struct {
	ID    string `json:"id" gorm:"primary_key" sql:"type:uuid;default:uuid_generate_v4()"`
	Day   *Day   `json:"day" gorm:"ForeignKey:DayID"`
	DayID string `json:"dayID" gorm:"type:uuid REFERENCES days(id);column:day_id"`
	Vow   string `json:"vow"`
	Plus  string `json:"plus"`
	Minus string `json:"minus"`
	Todo  string `json:"todo"`
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
	db.AutoMigrate(&Day{})
	db.AutoMigrate(&Entry{})

	user := &User{Name: "admin", Password: "ballet"}
	if db.NewRecord(user) {
		db.Create(user)
	}

	entries := []Entry{}
	db.Find(&entries)
	difference := 9 - len(entries)
	for i := 0; i < difference; i++ {
		day := &Day{
			Label: fmt.Sprintf("2018.01.0%d", (i + 1)),
		}
		db.Create(day)

		for j := 0; j < 4; j++ {
			entry := &Entry{
				DayID: day.ID,
				Vow:   "Protect life",
				Plus:  "I did well when...",
				Minus: "I need to improve on...",
				Todo:  "Spend two minutes each hour thinking about the nature of things",
			}
			db.Create(entry)
		}
	}

	return db
}
