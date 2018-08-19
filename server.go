package main

import (
	//"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type DayPayload struct {
	ID      string  `json:"id"`
	Label   string  `json:"label"`
	Entries []Entry `json:"entries"`
}

func setupRouter(db *gorm.DB) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:7000"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"Origin"},
	}))

	r.GET("/days/:dayId", func(c *gin.Context) {
		day := &Day{}
		db.Where("Label = ?", c.Param("dayId")).First(&day)
		entries := []Entry{}
		db.Where("day_id = ?", day.ID).Find(&entries)
		payload := &DayPayload{
			ID:      day.ID,
			Label:   day.Label,
			Entries: entries,
		}
		c.JSON(200, payload)
	})

	r.GET("/days/", func(c *gin.Context) {
		days := []Day{}
		db.Find(&days)
		c.JSON(200, days)
	})

	return r
}

func main() {
	db := setupDB()
	defer db.Close()

	r := setupRouter(db)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
