package main

import (
	//"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

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
		entry := &Entry{}
		db.Where("Day = ?", c.Param("dayId")).First(&entry)
		c.JSON(200, entry)
	})

	r.GET("/days/", func(c *gin.Context) {
		entries := []Entry{}
		db.Find(&entries)
		c.JSON(200, entries)
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
