package main

import (
	//"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		var dbuser User
		db.First(&dbuser, "name = ?", user)

		if dbuser.ID != 0 {
			c.JSON(200, gin.H{"password": dbuser.Password})
		} else {
			c.JSON(404, gin.H{"status": "not found"})
		}

	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		//user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		//var json struct {
		//	Value string `json:"value" binding:"required"`
		//}

		//if c.Bind(&json) == nil {
		//DB[user] = json.Value
		//c.JSON(200, gin.H{"status": "ok"})
		//}
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
