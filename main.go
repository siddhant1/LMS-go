package main

import (
	"ilearn-backend/db"
	"ilearn-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := db.Setup()

	router.Use(CORS())

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	routes.TeacherRoute(router)

	router.Run(":9000")
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
